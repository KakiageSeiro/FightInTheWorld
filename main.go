package main

import (
	"./logic"
	"fmt"
)

func main() {
	//reverse()
	//powersOf2()
	//sqrt()
	dupulicationTask()
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
