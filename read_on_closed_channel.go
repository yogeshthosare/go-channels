//go run What_will_happen_if_we_read_on_a_closed_channel.go
package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan bool)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case result, _ := <-a:
				fmt.Println(result)
				continue
			default:
				fmt.Println("here")
				continue
			}
		}
	}()
	a <- true
	time.Sleep(5 * time.Second)
	close(a)
	time.Sleep(20 * time.Second)
}

/*
  true
  here
  here
  here
  here
  here
  false
  false
  false
  false
  false
  false
  ________________________________*/
