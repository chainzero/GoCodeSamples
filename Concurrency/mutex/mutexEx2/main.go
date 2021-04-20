package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var i = 0

func main() {
	for ii := 0; ii < 100; ii++ {
		go Increment(&i)
		go Decrement(&i)
	}

	final := GetValue(&i)
	fmt.Println(*final)
}

func Increment(i *int) {
	mu.Lock()
	*i++
	mu.Unlock()
}

func Decrement(i *int) {
	mu.Lock()
	*i--
	mu.Unlock()
}

func GetValue(i *int) *int {
	mu.Lock()
	v := i
	mu.Unlock()
	return v
}
