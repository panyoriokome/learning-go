package main

import "fmt"

func main() {
	var array = [3]int{10, 20, 30}

	var slice = []int{10, 20, 30}
	fmt.Println(array)
	fmt.Printf("%T\n", array)
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	// Multidimensional slices
	// 定義方法がよくわからない...
	// var multi [][]int{10, 20, 30}
	// fmt.Println(multi)

	var array_ [3]int
	var slice_ []int
	fmt.Println(array_)
	fmt.Printf("%T\n", array_)
	fmt.Println(slice_)
	fmt.Printf("%T\n", slice_)

}
