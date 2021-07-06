package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0は割り算に使えないよ")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	result2, remainder2, err2 := divAndRemainder(5, 0)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Println(result2, remainder2)
}
