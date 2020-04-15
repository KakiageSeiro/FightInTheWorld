package chapter1

import "fmt"

type str struct {
	charSet  int32
	count    int
	hitCount int
}

func Permutation(array1, array2 string) bool {

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
