package inseong

import "fmt"

type cell2 struct {
	i, j     int
	canWrite bool
}

func (c *cell2) up()    { c.i-- }
func (c *cell2) down()  { c.i++ }
func (c *cell2) left()  { c.j-- }
func (c *cell2) right() { c.j++ }

func Ps02() {
	var N int
	fmt.Print("입력 : ")
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}
	arr := makeArr2(N)
	printArr2(arr)
}

func makeArr2(N int) [][]int {
	arr := make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = append(arr[i], make([]int, N)...)
	}

	num := 1
	c := &cell2{0, 0, true}
	for i := 0; i < N-1; i++ {
		num = writeNum(num, arr, c)
		c.right()
	}
	num = writeNum(num, arr, c)

	dirSwitch := true

	for i := N - 1; i > 0; i-- {
		if dirSwitch {
			for j := 0; j < i; j++ {
				c.down()
				num = writeNum(num, arr, c)
			}

			for j := 0; j < i; j++ {
				c.left()
				num = writeNum(num, arr, c)
			}

		} else {
			for j := 0; j < i; j++ {
				c.up()
				num = writeNum(num, arr, c)
			}

			for j := 0; j < i; j++ {
				c.right()
				num = writeNum(num, arr, c)
			}
		}
		dirSwitch = !dirSwitch
	}

	return arr
}

func printArr2(arr [][]int) {
	d := "%02d "
	if len(arr) > 9 {
		d = "%03d "
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf(d, arr[i][j])
		}
		fmt.Println()
	}
}

func writeNum(num int, arr [][]int, c *cell2) int {
	if c.canWrite {
		arr[c.i][c.j] = num
		num++
	} else {
		arr[c.i][c.j] = 0
	}
	c.canWrite = !c.canWrite
	return num
}
