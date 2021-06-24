package main

import "fmt"

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type myFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func myFunc(opts myFuncOpts) {
	fmt.Println("aaaa")
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(div(10, 2)) // 5
	fmt.Println(div(10, 0)) // 0
	fmt.Println(div(10, 3)) // 3

	myFunc(myFuncOpts{
		LastName: "John",
		Age:      20,
	})

	myFunc(myFuncOpts{
		LastName:  "John",
		FirstName: "Kane",
	})

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4))
}
