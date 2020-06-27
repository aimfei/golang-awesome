package base

import "fmt"

func main() {
	fmt.Print("a")
	var a = func(a int32, b int32) int32 {
		return a + b
	}(3, 4)
	fmt.Print(a)
}
