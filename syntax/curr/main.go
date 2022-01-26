package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("do 1")
		waitGroup.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("do 2")
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println("thread run finished")
}

func addValue(ch chan bool, value bool) {
	ch <- value
}
