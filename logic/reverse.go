package logic

import "fmt"

func Reverse(slise []int) {
	//引数のスライスのサイズの半分のみループ
	//O(N)
	for i := 0; i < len(slise)/2; i++ {
		other := len(slise) - i - 1
		temp := slise[i]
		slise[i] = slise[other]
		slise[other] = temp
	}

	fmt.Println(slise)
}
