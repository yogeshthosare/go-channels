//________________________________
//What_will_happen_if_we_write_on_a_closed_channel.go

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
			case result, ok := <-a:
				if !ok {
					fmt.Println(result)
					//here channel is closed and we are trying to write
					a <- true
				}
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
  Output -

  true
  here
  here
  here
  here
  false
  panic: send on closed channel

  goroutine 5 [running]:
  main.main.func1(0xc000058060)
          /home/yogesh/Go/src/goroutine_lim_patt/What_will_happen_if_we_write_on_a_closed_channel.go:19 +0x11c
  created by main.main
          /home/yogesh/Go/src/goroutine_lim_patt/What_will_happen_if_we_write_on_a_closed_channel.go:11 +0x5c
  ________________________________
*/
