package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

var sema = make(chan struct{}, 1)
var balance = 0

func main() {
	ch := make(chan int)
	go createMessage(ch, rand.Int())
	go receMessage(ch)
	<-ch

}

func createMessage(ch chan int, number int) {
	fmt.Printf("创建数据=%d\n", number)
	ch <- number
}
func receMessage(ch chan int) {
	message := <-ch
	fmt.Printf("接受数据=%d\n", message)
}

func deposit(amount int) int {
	sema <- struct{}{}
	mu.Lock()
	balance += amount
	<-sema
	mu.Unlock()
	rw.Lock()
	return balance
}

var (
	mu sync.Mutex
	rw sync.RWMutex
)

func newCtx() context.Context {
	return context.Background()
}
