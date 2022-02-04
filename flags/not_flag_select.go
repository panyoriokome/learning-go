package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Arg(0), flag.Arg(1))
}
