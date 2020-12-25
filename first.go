package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	sign := make(chan struct{}, 1)
	go func() {
		var e int
		ok := true
		for {
			select {
			case e, ok = <-ch:
				if !ok {
					break
				}
				fmt.Printf("Recieve a int from channel ch: %d\n", e)
			case ok = <-func() chan bool {
				ch := make(chan bool, 1)
				go func(ch chan bool) {
					time.Sleep(time.Second)
					ch <- false
				}(ch)
				return ch
			}():
				fmt.Println("Time out.")
				break
			}
			if !ok {
				break
			}
		}
		sign <- struct{}{}
	}()

	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(10 * time.Millisecond)
	}

	<-sign
}
