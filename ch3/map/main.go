package main

import "fmt"

func main() {
	// varを使った定義
	var nilMap map[string]int

	fmt.Println(nilMap) // map[]
	// mapのzero valueは0
	fmt.Println(len(nilMap)) // 0

	// :=でmap literalを使える
	total := map[string]int{}
	fmt.Println(total) // map[]

	//
	teams := map[string][]string{
		"Blue":  []string{"James", "Ramos"},
		"Red":   []string{"Nacho", "Modric"},
		"Green": []string{"Zidane", "Kovacic"},
	}
	// mapをすべて出力
	fmt.Println(teams) // map[Blue:[James Ramos] Green:[Zidane Kovacic] Red:[Nacho Modric]]
	// 指定したキーのみ出力
	fmt.Println(teams["Blue"]) // [James Ramos]
	// 指定したキーで取り出したmapの値から、インデックスを指定
	fmt.Println(teams["Blue"][0]) // James

	// makeの指定
	events := make(map[int][]string, 10)
	fmt.Println(events) // map[]

	// Reading and Writing
	totalSample := map[string]int{}
	totalSample["Blue"] = 1
	totalSample["Red"] = 2
	fmt.Println(totalSample["Blue"]) // 1
	fmt.Println(totalSample)
	fmt.Println(totalSample["Green"]) // 代入されていないため0になる
	totalSample["Green"]++
	fmt.Println(totalSample["Green"]) // 1
	totalSample["Red"] = 3
	fmt.Println(totalSample["Red"]) // 3

	// comma ok idiom
	m := map[string]int{
		"blue": 2,
		"red":  0,
	}
	v, ok := m["blue"]
	fmt.Println(v, ok) // 2 true

	v, ok = m["red"]
	fmt.Println(v, ok) // 0 true

	v, ok = m["green"]
	fmt.Println(v, ok) // 0 true

	// Deleting from Maps
	newMap := map[string]int{
		"blue": 2,
		"red":  1,
	}
	delete(newMap, "blue")
	fmt.Println(newMap)     // map[red:1]
	delete(newMap, "green") // 存在しないキーのため何も起こらない
	fmt.Println(newMap)     // map[red:1]

	// set in go
	intSet := map[int]bool{}
	vals := []int{5, 10, 3, 4, 5, 6, 7, 8}
	//
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println((len(vals)), len(intSet)) // 8 7
	fmt.Println(intSet[5])                // true
	fmt.Println(intSet[500])              // false
	if intSet[100] {
		fmt.Println("100がsetの中にある")
	}

}
