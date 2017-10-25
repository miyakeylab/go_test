/*********************************/
/* 課題1-4						 */
/*********************************/
package kadai

import "fmt"

type Stringer interface {
	String() string
}
//8進数型定義
type MyHex int
//16進数文字列表記
func (h MyHex) String() string {
	return fmt.Sprintf("%dの16進数：%x",int(h), int(h))
}
func Kadai_1_4(){

	var stringer Stringer
	// int型の100をHex型にキャストし，fmt.Stringer型の変数に代入している
	stringer = MyOct(100)
	TypeCheck(stringer)
	fmt.Println(stringer)

	stringer = Person {
		Name: "Myk",
		Age: 33,
		Loc:"熊本",
	}
	TypeCheck(stringer)
	fmt.Println(stringer)

	stringer = MyMap{"x":1}
	TypeCheck(stringer)
	fmt.Println(stringer)

	stringer = Alphabet{"a","b","c",}
	TypeCheck(stringer)
	fmt.Println(stringer)

	stringer = MyHex(100)
	TypeCheck(stringer)
	fmt.Println(stringer)

}

func TypeCheck(str Stringer){

	switch str.(type) {
	case MyOct://8進数
		fmt.Println("type:MyOct")
	case Person://構造体(人データ型)
		fmt.Println("type:Person")
	case MyMap://MAP型
		fmt.Println("type:MyMap")
	case Alphabet:// スライス(アルファベット)
		fmt.Println("type:Alphabet")
	default:
		fmt.Println("type:default")
	}

}