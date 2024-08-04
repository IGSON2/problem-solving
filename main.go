package main

import (
	"fmt"
	miningminerals "problem-solving/greedy/mining-minerals"
)

func main() {
	picks := []int{1, 3, 2}
	minerals := []string{"diamond", "diamond", "diamond", "iron", "iron", "diamond", "iron", "stone"}
	fmt.Println(miningminerals.MiningMinerals(picks, minerals))
}
