/*********************************/
/* 課題1-5						 */
/*********************************/
package kadai

import (
	"fmt"
	"errors"
)

type error interface {
	Error() string
}

type NoStringer interface {
}

func Kadai_1_5() {

	var s NoStringer

	if n, err := ToStringer(s); err != nil {
			fmt.Println("ErrorMessage:", err)
		} else {
			fmt.Println(n)
		}

	// Stringerあり
	s = MyOct(100)
	if n, err := ToStringer(s); err != nil {
		fmt.Println("ErrorMessage:", err)
	} else {
		fmt.Println(n)
	}
}

func ToStringer(v interface{})(Stringer,error){

	var err error = nil

	 n, ok := v.(Stringer)

	if ok == false{
		//err = fmt.Errorf("This type is not Stringer")
		err = errors.New("This type is not Stringer")
	}

	return n,err
}