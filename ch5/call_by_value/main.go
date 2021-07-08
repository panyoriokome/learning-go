package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	m[4] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	p := person{}
	i := 2
	s := "hello"
	// 値の変更に失敗する
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 hello {0 }

	// mapの場合
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m) // map[2:hello 3:goodbye 4:goodbye]

	// sliceの場合
	s2 := []int{1, 2, 3}
	modSlice(s2)
	fmt.Println(s2) // [2 4 6]
}
