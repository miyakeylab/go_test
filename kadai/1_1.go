/*********************************/
/* 課題1-1						 */
/*********************************/
package kadai

import "fmt"

/**
 * 課題1-1
 */
func Kadai_1_1(){

	// 1-1
	for i := 1; i <= 100; i++{

		if i % 2 == 0{
			fmt.Println(i,"-偶数")
		}else {
			fmt.Println(i, "-奇数")
		}
	}
}
