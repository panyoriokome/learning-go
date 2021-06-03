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

	fmt.Println("lenやappend等のBuilt in function")
	//fmt.Println(slice_ == slice_)

	fmt.Println(len(slice_))

	// 1つの要素を追加
	slice_ = append(slice_, 10)
	fmt.Println(slice_) // [10]

	// 複数要素を追加
	slice_ = append(slice_, 20, 30, 40)
	fmt.Println(slice_) // [10 20 30 40]

	var y = []int{50, 60, 70}
	slice_ = append(slice_, y...)
	fmt.Println(slice_) // [10 20 30 40 50 60 70]

	// Goはcall by valueの言語であるため、返り値を変数に代入しないとエラーになる
	// append(slice_, y...) // append(slice_, y...) evaluated but not used

	fmt.Println("Capacity")

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
