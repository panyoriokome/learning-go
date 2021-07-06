package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("0は割り算に使えないよ")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder2(numerator, denominator int) (result int, remainder int, err error) {
	// ここで変数にアサイン
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("0は割り算に使えないよ")
	}
	// 定義した変数名とは別の変数で返す
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
