package main

import "fmt"

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
}

func addTo(base int, vals ...int) []int {
	// 受け取った引数の数だけのArrayを作る
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
