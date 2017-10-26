package kadai

import (
	"fmt"
	"flag"

)

func Kadai_1_6(n bool) {

	flag.Parse()

	fmt.Println(flag.Args())
	var form string = "%s"

	if n == true{
		form = "%s\n"
	}

	fmt.Println(fmt.Scanf(form,flag.Args()))
	//MerucariCreate()
}

