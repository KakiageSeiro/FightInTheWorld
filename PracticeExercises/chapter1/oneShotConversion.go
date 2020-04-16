package chapter1

/*
1.5 一発変換
文字列に対して1文字だけの「挿入」「削除」「置き換え」のどれか1操作のみで、
もう片方の文字列にできるかを判定する
*/

func OneShotConversion(before, after string) bool {

	if len(before) == len(after) {
		//変換
		return convert2(before, after)
	} else if len(before) > len(after) {
		//削除
		return delete(before, after)
	} else if len(before) < len(after) {
		//挿入
		return insert(before, after)
	}

	return false
}

func insert(before, after string) bool {
	if len(before) != len(after)-1 {
		return false
	}

	beforeSlice := stringToSlice(before)
	afterSlice := stringToSlice(after)

	for i := range beforeSlice {

		//beforeとafterで違う文字がある場合、その文字をbeforeに挿入した結果同じ文字になる
		//違う文字を見つけたらafterの次の文字とも比較して、一致しない（1文字削除しても文字が違う）場合はfalseとする
		if beforeSlice[i] != afterSlice[i] {
			if beforeSlice[i] != afterSlice[i+1] {
				return false
			}
		}
	}

	return true
}

func delete(before, after string) bool {

	if len(before) != len(after)+1 {
		return false
	}

	beforeSlice := stringToSlice(before)
	afterSlice := stringToSlice(after)

	for i := range afterSlice {

		//beforeとafterで違う文字ある場合、その文字を削除した結果がafterになる
		//違う文字を見つけたらbeforeの次の文字とも比較して、一致しない（1文字削除しても文字が違う）場合はfalseとする
		if beforeSlice[i] != afterSlice[i] {
			if beforeSlice[i+1] != afterSlice[i] {
				return false
			}
		}
	}

	return true
}

//時間 O(n)
//空間 O(1)
func convert(before, after string) bool {
	//文字数は同じ前提
	if len(before) != len(after) {
		return false
	}

	beforeSlice := stringToSlice(before)
	afterSlice := stringToSlice(after)

	targetCount := 0
	for i := range beforeSlice {
		if beforeSlice[i] != afterSlice[i] {
			targetCount++
		}
	}

	//一文字だけ変換してafterにならない場合
	if targetCount != 1 {
		return false
	}

	return true
}

//計算量表記は変わらないが、変換しなければならない文字が2つ見つかった時点で処理を終了する分、計算削減したVer
//時間 O(n)
//空間 O(1)
func convert2(before, after string) bool {
	//文字数は同じ前提
	if len(before) != len(after) {
		return false
	}

	//同じ文字列の場合は変換する必要がない
	if before == after {
		return false
	}

	beforeSlice := stringToSlice(before)
	afterSlice := stringToSlice(after)

	targetCount := 0
	for i := range beforeSlice {
		if beforeSlice[i] != afterSlice[i] {
			targetCount++

			//一文字だけ変換してafterにならない場合
			if targetCount > 1 {
				return false
			}
		}
	}

	return true
}

func stringToSlice(s string) []int32 {
	//変換しなければならない文字数を数える
	result := []int32{}
	for _, v := range s {
		result = append(result, v)
	}
	return result
}
