package main

import "fmt"

func main() {
	ch := make(chan int)

	//send
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	//receive
	for v := range ch {
		fmt.Println(v)
	}
}
