package main

import (
	"./kadai"
	//"fmt"
	"flag"
)
//var n bool
//func init() {
//	// ポインタを指定して設定を予約
//	flag.BoolVar(&n, "n", false, "改行フラグ") }
//func main() {
//	fmt.Println("***********1-1**************")
//	kadai.Kadai_1_1()
//	fmt.Println("***********1-2**************")
//	kadai.Kadai_1_2()
//	fmt.Println("***********1-3**************")
//	kadai.Kadai_1_3()
//	fmt.Println("***********1-4**************")
//	kadai.Kadai_1_4()
//	fmt.Println("***********1-5**************")
//	kadai.Kadai_1_5()
//	fmt.Println("***********1-6**************")
//	kadai.Kadai_1_6(n)
//}

func main() {
	flag.Parse()
	kadai.Syukudai_1()
}