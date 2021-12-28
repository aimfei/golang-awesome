package sli

import "fmt"

func Fun1() {
	months := [...]string{1: "jan", 2: "feb", 12: "dec"}
	fmt.Println(months[0:2])
	fmt.Printf("size=%d\n", len(months))
	fmt.Printf("cap=%d\n", cap(months))
}
