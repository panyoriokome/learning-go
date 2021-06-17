package main

import "fmt"

func main() {
	// for-rangeの基本形(sliceに対して)
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

	// indexが不要な場合は_をつかう
	vals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range vals2 {
		fmt.Println(v)
	}
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12

	// 片方のみを使うケース
	// valueが不要な場合は変数自体を定義しない
	vals3 := []int{2, 4, 6, 8, 10, 12}
	for i := range vals3 {
		fmt.Println(i)
	}

	// mapでキーのみを使うケース
	names := map[string]bool{"John": true, "Tom": true, "Sam": true}
	for k := range names {
		fmt.Println(k)
	}

	// mapに対するfor文。3回同じ処理を繰り返す
	m := map[string]int{
		"a": 1,
		"b": 3,
		"c": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	// Loop 0
	// a 1
	// b 3
	// c 2
	// Loop 1
	// a 1
	// b 3
	// c 2
	// Loop 2
	// c 2
	// a 1
	// b 3

	// Debug用途のため、formatting functionでは常にキーの昇順で出力される
	for i := 0; i < 10; i++ {
		fmt.Println(m)
	}

	// strings
	samples := []string{"hello", "apple_n!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	// 0 104 h
	// 1 101 e
	// 2 108 l
	// 3 108 l
	// 4 111 o

	// 0 97 a
	// 1 112 p
	// 2 112 p
	// 3 108 l
	// 4 101 e
	// 5 95 _
	// 6 110 n
	// 7 33 !

	// for-rangeでアクセスする値はコピー
	vals4 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range vals {
		v *= 2
	}
	fmt.Println(vals4)

}
