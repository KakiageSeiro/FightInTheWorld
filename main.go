package main

import (
	"./PracticeExercises/chapter1"
	"./logic"
	"fmt"
)

//今着手している問題をここからキックする
//すでに完了した問題はコメントアウトしておく
func main() {
	println("#############################出力#################################")

	//■9章「問題」より前の章で登場した問題やサンプルコード■
	//reverse()
	//powersOf2()
	//sqrt()
	//dupulicationTask()
	//logic.UnnecessaryTask1()

	//■9章「問題」
	//isUniqueString()
	//permutation()
	//palindrome()
	//oneShotConversion()

	println(chapter1.Compression("aabcccccaaa"))
	println(chapter1.Compression("aaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbb"))
	println(chapter1.Compression("aabbccdd"))
	println(chapter1.Compression("aaBBccDD"))
	println("### 圧縮前の文字列のほうが長いパターン")
	println(chapter1.Compression("aabccdd"))
	println(chapter1.Compression("a"))
	println(chapter1.Compression("abcd"))
	println(chapter1.Compression("aaBbccDD"))

}

func oneShotConversion() {
	println("### 1文字削除で同じ文字列になる")
	println(chapter1.OneShotConversion("pale", "ple"))
	println(chapter1.OneShotConversion("pales", "pale"))
	println("### 1文字挿入で同じ文字列になる")
	println(chapter1.OneShotConversion("ple", "pale"))
	println(chapter1.OneShotConversion("ale", "pale"))
	println("### 1文字変換で同じ文字列になる")
	println(chapter1.OneShotConversion("pale", "bale"))
	println(chapter1.OneShotConversion("pale", "aale"))

	println("######################## ここからfalse ##########################")

	println("### 1文字削除で同じ文字列になる")
	println(chapter1.OneShotConversion("pale", "paa"))
	println(chapter1.OneShotConversion("pale", "pa"))
	println(chapter1.OneShotConversion("pale", ""))

	println("### 1文字挿入で同じ文字列になる")
	println(chapter1.OneShotConversion("pal", "palll"))
	println(chapter1.OneShotConversion("pale", "palaa"))
	println(chapter1.OneShotConversion("pale", "aaaaaaaaaaaaaaaaa"))

	println("### 1文字変換で同じ文字列になる")
	println(chapter1.OneShotConversion("pale", "bake"))
	println(chapter1.OneShotConversion("pale", "aaaa"))
	println(chapter1.OneShotConversion("pale", "pppe"))
}

func palindrome() {
	println(chapter1.Palindrome("a"))
	println(chapter1.Palindrome("aa aa aa"))
	println(chapter1.Palindrome("abcd cba"))
	println(chapter1.Palindrome("abcdd cba"))
	println(chapter1.Palindrome("abcd d cba"))
	println(chapter1.Palindrome("abcd dcba"))
	println(chapter1.Palindrome("aabbccddee"))
	println(chapter1.Palindrome("aabbccdde"))
	println(chapter1.Palindrome("a   dd ee ff"))
	println(chapter1.Palindrome("aaBBcc"))
	println(chapter1.Palindrome("aaBbBcc"))
	println(chapter1.Palindrome("Tact Coa"))
	println(chapter1.Palindrome("tact coa"))
	println(chapter1.Palindrome("aaBbBcc B"))
	println("###########ここからfalse#############")
	println(chapter1.Palindrome("ab"))
	println(chapter1.Palindrome("ab   dd ee ff"))
	println(chapter1.Palindrome("abcde edcba g h"))
}

func permutation() {
	println(chapter1.Permutation("aaabb", "aaabb"))
	println(chapter1.Permutation("aaabb", "babaa"))
	println(chapter1.Permutation("abbaa", "babaa"))
	println(chapter1.Permutation("aaabB", "aaabB"))

	println("###ここからfalse###")

	println(chapter1.Permutation("aaabb", "bbb"))
	println(chapter1.Permutation("aaabb", "aaabbb"))
	println(chapter1.Permutation("aaabb", "aacbb"))
	println(chapter1.Permutation("aaabb", "aaab b"))
	println(chapter1.Permutation("aaabb", "aaabB"))
}

func isUniqueString() {
	str1 := "abcde"
	str2 := "abcdef"
	str3 := "abcdefg"

	str4 := "aacbde"
	str5 := "abcdea"

	//println(chapter1.IsUniqueString1(str1))
	//println(chapter1.IsUniqueString1(str2))
	//println(chapter1.IsUniqueString1(str3))
	//
	//println(chapter1.IsUniqueString1(str4))
	//println(chapter1.IsUniqueString1(str5))

	//println(chapter1.IsUniqueString2(str1))
	//println(chapter1.IsUniqueString2(str2))
	//println(chapter1.IsUniqueString2(str3))
	//
	//println(chapter1.IsUniqueString2(str4))
	//println(chapter1.IsUniqueString2(str5))

	println(chapter1.IsUniqueString3(str1))
	println(chapter1.IsUniqueString3(str2))
	println(chapter1.IsUniqueString3(str3))

	println(chapter1.IsUniqueString3(str4))
	println(chapter1.IsUniqueString3(str5))
}

func dupulicationTask() {
	//logic.DuplicationTask1()
	logic.DuplicationTask2()
	//logic.DuplicationTask3()
}

func sqrt() {
	fmt.Println(logic.Sqrt(9))
	fmt.Println(logic.Sqrt(10))
	fmt.Println(logic.Sqrt(100))
}

func powersOf2() {
	//nの大きさが２倍になったタイミングで再帰呼び出し回数が増加していることがわかる
	logic.PowersOf2(31) //1,2,4,8,16の5つが出力される
	fmt.Println("■■■")
	logic.PowersOf2(32) //1,2,4,8,16,32の6つが出力される
	fmt.Println("■■■")
	logic.PowersOf2(63) //1,2,4,8,16,32の6つが出力される
	fmt.Println("■■■")
	logic.PowersOf2(64) //1,2,4,8,16,32,64の7つが出力される
}

func reverse() {
	ints := make([]int, 100)
	iota := 0
	for i := range ints {
		iota = iota + 1
		ints[i] = iota
	}
	fmt.Println(ints)

	logic.Reverse(ints)
}
