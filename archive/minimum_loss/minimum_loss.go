package minimumloss

import "sort"

type priceAndYear struct {
	year  int
	price int64
}

func minimumLoss(price []int64) int64 {
	var results []int64
	pny := make([]priceAndYear, 0)
	for i, p := range price {
		pny = append(pny, priceAndYear{i, p})
	}
	sort.Slice(pny, func(i, j int) bool {
		return pny[i].price < pny[j].price
	})
	for i := 0; i < len(pny)-1; i++ {
		if pny[i+1].year < pny[i].year {
			results = append(results, int64(pny[i+1].price-pny[i].price))
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results[0]
}
