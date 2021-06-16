package main

import "fmt"

func main() {
	vals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range vals {
		fmt.Println(i, v)
	}
	// 0 2
	// 1 4
	// 2 6
	// 3 8
	// 4 10
	// 5 12
}
