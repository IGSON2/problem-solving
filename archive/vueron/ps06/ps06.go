package ps06

type dot struct {
	x, y int
	dist float64
}

func ps06(arr [][2]int) int {
	// polygon := make([]dot, 0)
	dots := make([]dot, 0)
	for i := 0; i < len(arr); i++ {
		d := dot{
			x: arr[i][0],
			y: arr[i][1],
		}
		d.dist = sqrt(d.x*d.x + d.y*d.y)
		insertDot(&dots, d)
	}
	return 0
}

func sqrt(x int) float64 {
	if x <= 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	n := 0.0
	for i := 1; i <= int(x/2)+1; i++ {
		if i*i == x {
			return float64(i)
		}

		if i*i > x {
			n = float64(i)
			break
		}
	}

	decimal := 0.00001
	minDiff := float64(x)
	min := float64(n-1) + decimal

	for i := n - 1; i < n; i += decimal {
		diff := float64(x) - i*i
		if diff > 0 && diff < minDiff {
			minDiff = diff
			min = i
		}

		if diff < 0 {
			break
		}
	}
	return abs(min)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func insertDot(dots *[]dot, d dot) {
	if len(*dots) == 0 {
		*dots = append(*dots, d)
		return
	}

	for i := 0; i < len(*dots); i++ {
		if (*dots)[i].dist < d.dist {
			*dots = append([]dot{d}, *dots...)
			return
		}
	}
	*dots = append(*dots, d)
}
