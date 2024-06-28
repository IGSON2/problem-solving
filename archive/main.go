package main

import (
	biggerisgreater "coding_test/bigger_is_greater"
	"coding_test/utils"
	"fmt"
)

func main() {
	testStr := utils.MakeRanString(5)
	resStr := biggerisgreater.BiggerIsGreater(testStr)
	fmt.Printf("Test string: %s\nResult string: %s\n", testStr, resStr)
}
