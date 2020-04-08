package logic

//平方根を求める
//O(sqrt(n))
func Sqrt(n int) int {
	//２乗の数(平方根）を順番に調べていく
	for guess := 1; guess*guess <= n; guess++ {
		if guess*guess == n {
			return guess
		}
	}
	//引数が２乗の数でない場合はエラー
	return -1
}
