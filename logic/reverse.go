package logic

import "fmt"

//文字列を反転させる
//O(N)
func Reverse(slise []int) {
	//引数のスライスのサイズの半分のみループ
	for i := 0; i < len(slise)/2; i++ {
		other := len(slise) - i - 1
		temp := slise[i]
		slise[i] = slise[other]
		slise[other] = temp
	}

	fmt.Println(slise)
}
