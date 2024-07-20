package inseong

import "fmt"

type cell struct {
	i, j int
}

func (c *cell) up()    { c.i-- }
func (c *cell) down()  { c.i++ }
func (c *cell) left()  { c.j-- }
func (c *cell) right() { c.j++ }

func Ps01() {
	var N int
	fmt.Print("입력 : ")
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}
	arr := makeArr(N)
	printArr(arr)
}

func makeArr(N int) [][]int {
	num := 1
	arr := make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = append(arr[i], make([]int, N)...)
	}

	for i := 0; i < N; i++ {
		arr[0][i] = num
		num++
	}

	c := cell{0, N - 1}
	dirSwitch := true

	for i := N - 1; i > 0; i-- {
		if dirSwitch {
			for j := 0; j < i; j++ {
				c.down()
				arr[c.i][c.j] = num
				num++
			}

			for j := 0; j < i; j++ {
				c.left()
				arr[c.i][c.j] = num
				num++
			}

		} else {
			for j := 0; j < i; j++ {
				c.up()
				arr[c.i][c.j] = num
				num++
			}

			for j := 0; j < i; j++ {
				c.right()
				arr[c.i][c.j] = num
				num++
			}
		}
		dirSwitch = !dirSwitch
	}

	return arr
}

func printArr(arr [][]int) {
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
