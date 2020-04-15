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
