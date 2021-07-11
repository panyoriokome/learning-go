package main

import "fmt"

func main() {
	x := 10
	// &はaddress operatorであり、指定した変数の値が格納されているメモリ上の場所を返す
	pointerToX := &x

	fmt.Println(pointerToX) // 0xc0000160d8

	// *はindirection operatorであり、pointerを受け取ってpointerが示す値を返す
	fmt.Println(*pointerToX) // 10

	// var z *int
	// fmt.Println(z == nil)
	// fmt.Println(*z) // panic: runtime error: invalid memory address or nil pointer dereference

	x2 := 15
	var pointerToX2 *int
	pointerToX2 = &x2

	fmt.Println(pointerToX2)  // 0xc000016110
	fmt.Println(*pointerToX2) // 15

	// x2 := &Foo{}
	// var y2 string
	// z2 := &y2
	// fmt.Println(z2)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: "Perry",
		LastName:   "Peterson",
	}
}
