package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{10, 20, 30}
	var sparse = [12]int{1, 5: 4, 6, 10: 100, 18}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(sparse)
}
