package main

func main() {

}

type BaseResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
