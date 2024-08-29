package carpet

type yellowArea struct {
	height, width int
}

func solution(brown int, yellow int) []int {
	yAreas := make([]yellowArea, 0)
	for i := 1; ; i++ {
		width := yellow / i
		if width < i {
			break
		}
		if yellow%i == 0 {
			yAreas = append(yAreas, yellowArea{height: i, width: width})
		}
	}

	for _, y := range yAreas {
		if 2*(y.width+2)+2*y.height == brown {
			return []int{y.width + 2, y.height + 2}
		}
	}
	panic("invalid")
}

func solution2(brown int, red int) []int {
	n := brown + red

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			w, h := i, n/i
			if 2*h+2*w == brown+4 {
				return []int{h, w}
			}
		}
	}

	panic("error")
}
