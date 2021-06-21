package main

import "fmt"

func main() {
	// for-rangeで記載
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v)
	}

	// Complete, C-style for
	evenVals2 := []int{2, 4, 6, 8, 10}
	for i := 1; i < len(evenVals2)-1; i++ {
		fmt.Println(i, evenVals2[i])
	}
}
