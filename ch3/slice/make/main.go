package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Println(x)

	// makeで定義したSliceにappendを使う例
	x = append(x, 10)
	fmt.Println(x) // [0 0 0 0 0 10]

	x[0] = 5
	fmt.Println(x) // [5 0 0 0 0 10]

}
