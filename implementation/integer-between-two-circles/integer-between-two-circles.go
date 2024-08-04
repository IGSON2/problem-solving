package integerbetweentwocircles

import (
	"fmt"
	"math"
)

func IntBetween2Circles(r1, r2 int) int64 {
	var cnt int64
	for x := r2; x > 0; x-- {
		y1 := 0.0
		if r1 < x {
			y1 = 0
		} else {
			y1 = math.Ceil(math.Sqrt(float64(r1*r1 - x*x)))
		}
		y2 := math.Floor(math.Sqrt(float64(r2*r2 - x*x)))
		fmt.Println(x, y1, y2)
		cnt += int64(y2) - int64(y1) + 1

	}
	return cnt * 4
}
