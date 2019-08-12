package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d", i)
	}
	wg.Done()
}
func main() {
	defer func() {
		fmt.Println("---------------exit")
	}()
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	go func() {

		for i := 1; i <= 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}
	}()
	go func(ch chan int) {
		time.Sleep(time.Second * 11)
		close(ch)
	}(c)
	wg.Wait()

}
