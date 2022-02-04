package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")
	fmt.Println(nFlag)
}
