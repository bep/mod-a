package main

import (
	"fmt"

	ma "github.com/bep/mod-a/p"
	mc "github.com/bep/mod-c/p"
)

func main() {
	fmt.Println("mod-a:", ma.Version())
	fmt.Println("mod-c:", mc.Version())
}

func ModA() string {
	return "ModA"
}
