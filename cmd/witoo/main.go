package main

import (
	"fmt"
	"log"

	"github.com/iwpnd/witoo"
)

func main() {
	name := "Ben"
	s, err := witoo.SayHiTo(name)
	if err != nil {
		log.Fatalf("saying hi to %s: %s", name, err)
	}
	fmt.Println(s)
}
