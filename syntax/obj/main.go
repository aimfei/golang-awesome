package main

import "fmt"

func main() {
	dog := Dog{Attr0: "123",
		Attr1: "123",
	}
	dog.Age = 11
	dog.Name = "sdfs"

	fmt.Println(dog)
}

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Dog struct {
	Animal
	Attr0 string `json:"attr_0"`
	Attr1 string `json:"attr_1"`
}
