package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	// エラーが出る
	// fmt.Println(person)

	var sakata person
	fmt.Println(sakata) // { 0 }

	yamada := person{}
	fmt.Println(yamada) // { 0 }

	// 定義方法
	// 1. キーを指定しないで定義
	john := person{
		"John",
		25,
		"dog",
	}
	fmt.Println(john) // {John 25 dog}

	// 2. キーを指定して定義
	ian := person{
		age:  29,
		name: "Bez",
	}
	fmt.Println(ian) // {Bez 29 }

	// フィールドにはドットでアクセスする
	fmt.Println(ian.name) // Bez

	// Anonymous Struct

	// 1. varでのAnonymous Struct
	var todo struct {
		id    int
		title string
	}

	todo.id = 1
	todo.title = "牛乳を買う"
	fmt.Println(todo) // {1 牛乳を買う}

	// 2 :=でのAnonymous Struct
	product := struct {
		id   int
		name string
	}{
		id:   1,
		name: "おいしい牛乳",
	}
	fmt.Println(product) // {1 おいしい牛乳}

}
