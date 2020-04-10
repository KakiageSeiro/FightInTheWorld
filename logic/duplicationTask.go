package logic

import (
	"fmt"
)

type cd struct {
	C int
	D int
}

//cube(a)+cube(b) == cube(c)+cube(d)がtrueとなる全ての組み合わせを試す
//a,bの組み合わせに対して、c,dの組み合わせが複数回登場する。登場するたびに冪乗計算をしているため、
//先行して計算結果をキャッシュしておく
//O(n^4)だけど、計算しなくなった分は減っている
func UnnecessaryTask1() {
	hitCount := 0 //cube(a)+cube(b) == cube(c)+cube(d)がtrueとなる全ての組み合わせの数
	n := 10       //a,b,c,dが取りうる最大数(最小値は1)

	//c,dの計算結果をMapに保存しておく(key:計算結果, value:cとdの組み合わせのslise)
	m := createCDMap(n)

	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			//c,dのマップから、a,bの計算結果と同じものを探す
			cdList, ok := m[cube(a)+cube(b)]
			if !ok {
				continue
			}

			//cube(a)+cube(b) == cube(c)+cube(d)がtureなものを出力
			for i := range cdList {
				hitCount++
				fmt.Println(a, b, cdList[i].C, cdList[i].D)
			}
		}
	}

	fmt.Println(hitCount)
}

func createCDMap(n int) map[int][]cd {
	m := make(map[int][]cd)
	for c := 1; c <= n; c++ {
		for d := 1; d <= n; d++ {
			result := cube(c) + cube(d)
			cdPair := cd{c, d}
			if len(m[result]) > 0 {
				m[result] = append(m[result], cdPair)
			} else {
				m[result] = []cd{cdPair}
			}
		}
	}
	return m
}
