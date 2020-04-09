package logic

import (
	"fmt"
	"math"
)

// 3乗の値を返す
func cube(x int) int {
	//return x * x * x
	return int(math.Pow(float64(x), 3))
}

//cube(a)+cube(b) == cube(c)+cube(d)がtrueとなる全ての組み合わせを試す
//O (n^4)
func DuplicationTask1() {
	hitCount := 0

	n := 10
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				for d := 1; d <= n; d++ {
					if cube(a)+cube(b) == cube(c)+cube(d) {
						hitCount++

						fmt.Printf("%d %d %d %d", a, b, c, d)
						fmt.Println()
					}
				}
			}
		}
	}

	fmt.Println(hitCount)
}

//cube(a)+cube(b) == cube(c)+cube(d)がtrueとなる全ての組み合わせを試す
//dのループ内にtrueとなる組み合わせを一つしかないため、ヒットした時点でループを抜ける
//O(n^4)だけど、少しは減っている
func DuplicationTask2() {
	hitCount := 0

	n := 10
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				for d := 1; d <= n; d++ {
					if cube(a)+cube(b) == cube(c)+cube(d) {
						hitCount++

						fmt.Printf("%d %d %d %d", a, b, c, d)
						fmt.Println()

						break //dのループを抜ける
					}
				}
			}
		}
	}

	fmt.Println(hitCount)
}

//cube(a)+cube(b) == cube(c)+cube(d)がtrueとなる全ての組み合わせを試す
//dのループ内にtrueとなる組み合わせを一つしかないため、計算によってdを求める
//TODO:■■■求めたかったけどPowで分数(1/3)を扱っているからか、うまく計算できてない・・・
//O(n^3) になるはず・・・
func DuplicationTask3() {
	hitCount := 0
	exponential := float64(1) / 3

	n := 10
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				//後続のifを満たすdの値を求める
				d := math.Pow(float64(cube(a)+cube(b)-cube(c)), exponential)

				//デバッグ用（後続のif判定）
				aaa := cube(a)+cube(b) == cube(c)+cube(int(d))
				fmt.Println(aaa)


				//dは計算によって求めたので、整数値であることを確認する条件も追加している
				if cube(a)+cube(b) == cube(c)+cube(int(d)) && 0 <= d && int(d) <= n {
					hitCount++

					fmt.Printf("%d %d %d %d", a, b, c, int(d))
					fmt.Println()
				}
			}
		}
	}

	fmt.Println(hitCount)
}
