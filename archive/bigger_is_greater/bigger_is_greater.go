package biggerisgreater

import "sort"

func BiggerIsGreater(w string) string {
	res := "no answer"
	rw := []rune(w)
	swapIdx := struct {
		src int
		dst int
	}{}

	for i := len(rw) - 2; i >= 0; i-- {
		gap := int(rune('z'))
		for j := i + 1; j < len(rw); j++ {
			if rw[j] > rw[i] {
				if newgap := int(rw[j]) - int(rw[i]); newgap < gap {
					gap = newgap
					swapIdx.src, swapIdx.dst = i, j
				}
			}
		}
		if gap < int(rune('z')) {
			break
		}
	}

	if swapIdx.src != 0 || swapIdx.dst != 0 {
		rw[swapIdx.src], rw[swapIdx.dst] = rw[swapIdx.dst], rw[swapIdx.src]
		sort.Slice(rw[swapIdx.src+1:], func(i, j int) bool {
			return rw[swapIdx.src+1+i] < rw[swapIdx.src+1+j]
		})
		res = string(rw)
	}

	return res
}
