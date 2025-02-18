package main

import (
	"fmt"

	"github.com/iwpnd/witoo"
)

func main() {
	name := "Ben"
	s := witoo.SayHiTo(name)
	fmt.Println(s)
}
