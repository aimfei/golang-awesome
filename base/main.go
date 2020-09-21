package base

import "fmt"

func add() {
	var a = func(a int32, b int32) int32 {
		return a + b
	}(3, 4)
	fmt.Println(a)
}
