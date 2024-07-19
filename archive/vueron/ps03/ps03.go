package ps03

import (
	"bufio"
	"fmt"
	"net"
)

const (
	addr = "localhost:8080"
)

func Ps03() {
	s := NewServer(addr)
	go s.Listen()
	fmt.Printf("server listening on %s.\n", addr)

	c := NewClient(addr)

	defer s.Close()
	defer c.conn.Close()

	ch := make(chan struct{})
	go func() {
		for {
			fmt.Print(c.Read())
			ch <- struct{}{}
		}
	}()

	for {
		var msg string
		if msg == "exit" {
			break
		}
		fmt.Scanln(&msg)
		c.Write(msg)
		<-ch
	}

}

type Server struct {
	net.Listener
}

func NewServer(addr string) *Server {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	return &Server{listener}

}

func (s *Server) Listen() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("new connection from %s\n", conn.RemoteAddr())

		go func(c net.Conn) {
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				msg := scanner.Text()
				s.Write(conn, fmt.Sprintf("server says : %s", msg))
			}
		}(conn)
	}
}

func (s *Server) Write(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg + "\n"))
	if err != nil {
		panic(err)
	}
}

type Client struct {
	conn net.Conn
}

func NewClient(addr string) *Client {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return &Client{conn}
}

func (c *Client) Read() string {
	reader := bufio.NewReader(c.conn)
	msg, _ := reader.ReadString('\n')
	return msg
}

func (c *Client) Write(msg string) {
	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		panic(err)
	}
}
