package chapter1

/*
1.2 順列チェック
2つの文字列の片方が、もう片方の並び替えであるかを判定する
*/

import "fmt"

type str struct {
	charSet  int32
	count    int
	hitCount int
}

//時間計算量 O(n)
//空間計算量 O(n)
func Permutation(array1, array2 string) bool {

	//array1よりarray2のほうが文字数が大きい場合、array1から作成したMAPであるstrsに存在する文字のみが、
	//array2に同数存在するかを検証してしまうため、"aaa"と"aaab"が並び替え文字列として判定されてしまう
	//そのため文字列の長さが違う時点でfalseにする
	if len(array1) != len(array2) {
		return false
	}

	strs := []str{}
	for _, v := range array1 {
		//すでに登場した文字の場合はカウントをインクリメント
		exist := false
		for i, s := range strs {
			if s.charSet == v {
				strs[i].count++
				exist = true
			}
		}

		//存在しなかった場合は新たに型を作成してスライスに追加
		if !exist {
			newStr := str{
				charSet:  v,
				count:    1,
				hitCount: 0,
			}
			strs = append(strs, newStr)
		}
	}
	fmt.Printf("%v\n", strs)

	//array2の文字がarray1に含まれている場合は、その文字のhitCountをインクリメント
	for _, v := range array2 {
		for i2, strElement := range strs {
			if v == strElement.charSet {
				strs[i2].hitCount++
			}
		}
	}

	//全ての文字のhitCountがcountと同じ数字の場合は、array1と2は並び替えである
	for _, s := range strs {
		if s.count != s.hitCount {
			return false
		}
	}

	return true
}
