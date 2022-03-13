package main

import (
	"fmt"
	"github.com/google/gofuzz"
)

func main() {
	//f := fuzz.New()
	//var myInt int
	//f.Fuzz(&myInt)
	//fmt.Println(myInt)

	f := fuzz.New().NilChance(0).NumElements(1, 1)
	var myMap map[string]string
	f.Fuzz(&myMap)
	fmt.Println(myMap)
}
