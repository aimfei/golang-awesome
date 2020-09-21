package main

import (
	"encoding/json"
	"fmt"
	"golang-awesome/base"
)

func main() {
	res := grammar.Success("张三")
	resjson, _ := json.Marshal(res)
	fmt.Print(string(resjson))
}
