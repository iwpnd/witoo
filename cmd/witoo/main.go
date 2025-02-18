package main

import (
	"fmt"

	"github.com/iwpnd/witoo/v2"
)

func main() {
	name := "Ben"
	s := witoo.SayHiTo(name)
	fmt.Println(s)
}
