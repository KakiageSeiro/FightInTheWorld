package logic

import "fmt"

//２の累乗（２を掛けた値を返す）
//O(log n)
func PowersOf2(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		fmt.Println(1)
		return 1
	} else {
		prev := PowersOf2(n / 2)
		curr := prev * 2
		fmt.Println(curr)
		return curr
	}
}
