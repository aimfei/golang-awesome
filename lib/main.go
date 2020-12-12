package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Environ())
	fmt.Println(os.TempDir())
	os.Create("/Users/flying/go/src/golang-awesome/1.sql")
}
