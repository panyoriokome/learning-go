package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// n := rand.Intn(10)
	// fmt.Println(n) // 常に1

	// if n == 0 {
	// 	fmt.Println("低すぎる")
	// } else if n > 5 {
	// 	fmt.Println("高すぎる")
	// } else {
	// 	fmt.Println("ちょうどいい") // 常に1なのでここが表示される
	// }

	if n := rand.Intn(10); n == 0 {
		fmt.Println("低すぎる", n) //
	} else if n > 5 {
		fmt.Println("高すぎる", n)
	} else {
		fmt.Println("ちょうどいい", n)
	}
	fmt.Println(n) // undefined: n
}
