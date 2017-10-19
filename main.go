package main

import (
	"fmt"
	//"os"
	//"path/filepath"
)

func main() {
	kadai_1_1()
	kadai_1_2()
}
/**
 * 課題1-1
 */
func kadai_1_1(){

	// 1-1
	for i := 1; i <= 100; i++{

		if i % 2 == 0{
			fmt.Println(i,"-偶数")
		}else {
			fmt.Println(i, "-奇数")
		}
	}
}
/**
 * 課題1-2
 */
func kadai_1_2(){

	// 1-2
	for i := 1; i <= 100; i++{
		switch{
		case isOdd(i):
			fmt.Println(i,"-奇数2")
		default:
			fmt.Println(i,"-偶数2")

		}
	}
}
/**
 * 奇数判定
 */
func isOdd(i int) (check bool){

	if i % 2 == 0{
		check = false
	}else {
		check = true
	}
	return
}
