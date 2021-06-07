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
	totalSample["Orcas"] = 1
	totalSample["Lions"] = 2
	fmt.Println(totalSample["Orcas"])
	fmt.Println(totalSample["Kittens"]) // 代入されていないため0になる
	totalSample["Kittens"]++
	fmt.Println(totalSample["Kittens"])
	totalSample["Lions"] = 3
	fmt.Println(totalSample["Lions"])

}
