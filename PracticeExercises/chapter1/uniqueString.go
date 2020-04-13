package chapter1

/*
1.1 重複のない文字列
文字列が全て固有であるかどうかを判定する
*/

//一文字ずつ確認していく
//時間計算量 O(n^2)
//空間計算量 O(1)
func IsUniqueString1(target string) bool {
	roopCount := 0
	for i, c := range target {
		for i2, c2 := range target {

			if i == i2 {
				//同じ位置の文字は比較しない
				continue
			}

			roopCount++
			if c == c2 {
				println("ループ回数:", roopCount)
				return false
			}
		}
	}

	println("ループ回数:", roopCount)

	return true
}

//重複する処理を省く
//時間計算量 O((n^2)/2)
//空間計算量 O(1)
func IsUniqueString2(target string) bool {
	roopCount := 0
	for i, c := range target {
		for i2, c2 := range target {

			//同じ位置の文字は比較しない
			if i == i2 {
				continue
			}

			//すでに確認したことがある組み合わせはスキップ
			if i > i2 {
				continue
			}

			roopCount++

			if c == c2 {
				println("ループ回数:", roopCount)
				return false
			}
		}
	}

	println("ループ回数:", roopCount)

	return true
}

//ハッシュテーブルから値を取得する計算量はO(1)であることを利用する(ハッシュ衝突は考慮しない)
//時間計算量 O(N)
//空間計算量 O(N)
func IsUniqueString3(target string) bool {
	//[文字コード]文字の位置
	charsetMap := make(map[int32]int)
	for i, c := range target {
		charsetMap[c] = i
	}

	for i, c := range target {
		//ハッシュテーブルに存在
		i2, ok := charsetMap[c]
		if ok {
			//同じ位置の文字は比較しない
			if i == i2 {
				continue
			}

			return false
		}
	}

	return true
}
