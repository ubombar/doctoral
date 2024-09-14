package main

import (
	"fmt"

	e "github.com/ubombar/doctoral/pkg/engine"
)

func main() {
	fmt.Println("Started")
	t := e.Or(e.Const("hello"), e.Const("world"))

	lin, err := e.Linearlize(t, true)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for _, v := range lin {
		fmt.Printf("%q ", v)
	}

	fmt.Println("")

	// m, err := regexp.MatchString("^\\W$", " ")

	// if err != nil {
	// 	return
	// }

	// fmt.Printf("m: %v\n", m)
}
