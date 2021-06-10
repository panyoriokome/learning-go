package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5         // shadowing variables
		fmt.Println(x) //5 shadowing variablesにアクセス
	}
	fmt.Println(x) // 10  Top-levelのxにアクセス

	a := 10
	if a > 5 {
		a, b := 5, 20
		fmt.Println(a, b) // 5 20
	}
	fmt.Println(a) // 10
}
