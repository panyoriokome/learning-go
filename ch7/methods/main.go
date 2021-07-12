package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	tom := Person{
		FirstName: "Tom",
		LastName:  "Jones",
		Age:       25,
	}

	fmt.Println(tom.String())
}
