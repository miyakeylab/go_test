/*********************************/
/* 課題1-3						 */
/*********************************/
package kadai

import (
	"fmt"
)

//8進数型定義
type MyOct int
//8進数文字列表記
func (h MyOct) String() string {
	return fmt.Sprintf("%dの8進数：%o",int(h), int(h))
}
//人データ型定義
type Person struct {
	Name string
	Age int
	Loc string
}
func (p Person) String() string {
	return fmt.Sprintf("名前：%s 年齢:%d 出身地：%s",p.Name, p.Age ,p.Loc)
}

// MAP型定義
type MyMap map[string]int
func (m MyMap) String() string {
	var rtn string

	// 存在チェック(あれば文字列に追加)
	if n,ok :=m["x"]; ok{
		rtn += fmt.Sprintf("key：x value:%d ",n)
	}
	if n,ok :=m["y"]; ok{
		rtn += fmt.Sprintf("key：y value:%d ",n)
	}
	if n,ok :=m["z"]; ok{
		rtn += fmt.Sprintf("key：z value:%d ",n)
	}
	return rtn
}

//スライス型
type  Alphabet []string
func (a Alphabet) String() string {
	// len,cap,内容を表示
	return fmt.Sprintf("len：%d cap:%d ",len(a),cap(a))
}


func Kadai_1_3() {

	// 8進数の型
	var oct MyOct = 100
	fmt.Println(oct.String())

	// 人データ型
	var per Person = Person {
							Name: "Myk",
							Age: 33,
							Loc:"熊本",
							}
	fmt.Println(per.String())

	// MAP型
	var tmpMap MyMap = MyMap{"x":1}
	fmt.Println(tmpMap.String())
	tmpMap["y"] = 2
	fmt.Println(tmpMap.String())
	tmpMap["z"] = 3
	fmt.Println(tmpMap.String())

	//スライス型
	var alpha Alphabet = Alphabet{"a","b","c",}
	fmt.Println(alpha.String())
	alpha = append(alpha, "d", "e")
	fmt.Println(alpha.String())
}
