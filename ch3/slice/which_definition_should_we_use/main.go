package main

import "fmt"

func main() {
	// もしnilのままなら...
	var data []int
	fmt.Println(data) // []

	if data == nil {
		fmt.Println("dataはnil") // 出力される=dataはnil
	}

	var x = []int{}
	fmt.Println(x)
	if x == nil {
		fmt.Println("dataはnil") // 出力されない=xはnilではない
	}

}
