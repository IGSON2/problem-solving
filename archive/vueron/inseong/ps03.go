package inseong

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	addr = "localhost:8080"
)

func Ps03() {
	ch := make(chan struct{})

	s := NewServer(addr)
	go s.Listen(ch)
	fmt.Printf("server listening on %s.\n", addr)

	c := NewClient(addr)

	defer s.Close()
	defer c.conn.Close()

	go func() {
		for {
			fmt.Print(c.Read())
			ch <- struct{}{}
		}
	}()

	<-ch
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
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

func (s *Server) Listen(ready chan<- struct{}) {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("new connection from %s\n", conn.RemoteAddr())

		ready <- struct{}{}

		go func(c net.Conn) {
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				msg := scanner.Text()
				_, err := conn.Write([]byte(fmt.Sprintf("server says : %s\n", msg)))
				if err != nil {
					panic(err)
				}
			}
		}(conn)
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
