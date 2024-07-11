package sortingnumber2

import (
	"fmt"
	"os"
)

func sortingNumber2() {
	f, err := os.Open("testcase")
	if err != nil {
		panic(err)
	}
	var n int
	fmt.Fscanln(f, &n)

	negArr := make([]bool, 1000001)
	nonNegArr := make([]bool, 1000001)
	for i := 0; i < n; i++ {
		var e int
		fmt.Fscanln(f, &e)
		if e < 0 {
			negArr[-e] = true
		} else {
			nonNegArr[e] = true
		}
	}

	for i := 1000000; i >= 0; i-- {
		if negArr[i] {
			fmt.Println(-i)
		}
	}
	for i, isExist := range nonNegArr {
		if isExist {
			fmt.Println(i)
		}
	}
}
