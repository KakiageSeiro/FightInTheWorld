package chapter1

import (
	"strconv"
	"strings"
)

/*
1.6 文字列圧縮
連続する文字は数字を使って表現する文字列作成する
"aabcccccaaa"を"a2b1c5a3"のように連続する文字数を追加した文字列を作成し
その文字列が元の文字列より短い場合は返却する
同じ、もしくは長い場合は元の文字列を返却する
*/

//時間 O(n)
//空間 O(1)
//
//※今回はやらなかったけど、大量の文字を読み込むことを想定する場合、
//圧縮後の文字数だけを先にカウントする処理(文字列結合はしない) を実施し
//その後下記の処理を実施する方法をとったほうが早いかも
//
//メリット
//1. strings.joinの中で配列サイズの変更(アロケーション)が行われている前提で、先に容量を宣言することができる
//2. 圧縮後文字数が判明した時点で処理を終了できることがある
//
//デメリット
//1. 引数の文字列サイズ回数のループが2回発生する
func Compression(s string) string {
	var beforeCharset int32 //直前のループでの文字コード
	beforeCharset = -1      //初期値として存在しない文字コードを設定
	continuousCount := 0    //同じ文字が連続している回数
	resultStrSlice := []string{}
	for i, v := range s {
		//初回ループ時は初期化処理のみ
		if i == 0 {
			continuousCount = 1
			beforeCharset = v

			continue
		}

		if beforeCharset == v {
			continuousCount++
		} else {
			//結果文字列に前回のループでの文字と、その登場回数を連結する
			resultStrSlice = append(resultStrSlice, string(beforeCharset)+strconv.Itoa(continuousCount))
			continuousCount = 1
			beforeCharset = v
		}
	}

	//最後の文字種を連結
	resultStrSlice = append(resultStrSlice, string(beforeCharset)+strconv.Itoa(continuousCount))
	result := strings.Join(resultStrSlice, "")

	//圧縮前の文字列のほうが長い場合は、圧縮前文字列をそのまま返却する
	if len(result) > len(s) {
		return s
	} else {
		return result
	}
}
