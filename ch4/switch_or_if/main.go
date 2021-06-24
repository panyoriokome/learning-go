package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("低すぎる", n)
	case n > 5:
		fmt.Println("高すぎる", n)
	default:
		fmt.Println("ちょうどいい", n)
	}

}
