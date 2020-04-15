package chapter1

/*
1.4 回文の順列
文字列が回文の順列である場合true
※文字を並び替えて回文になればOK
*/

import (
	"fmt"
	"strings"
)

//時間計算量 O(N)
//空間計算量 O(N)
func Palindrome(s string) bool {

	//空白の除去(全角とかはとりあえず考慮しない)
	s2 := strings.Replace(s, " ", "", -1)

	//[文字コード]文字の登場回数
	m := make(map[int32]uint)
	for _, v := range s2 {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	fmt.Printf("%v\n", m)

	//奇数回登場する文字種を数える
	oddCount := 0
	for _, count := range m {
		if count%2 != 0 {
			//奇数の場合はインクリメント
			oddCount++
		}
	}

	//奇数回登場する文字種は、2つ以上ある場合回文にできない
	if oddCount > 1 {
		return false
	} else {
		return true
	}
}
