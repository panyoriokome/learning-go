package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag int
	flag.IntVar(&nFlag, "n", 1234, "help message for flag n")
	flag.Parse()
	fmt.Println(nFlag)
}
