//________________________________
//What will happen if we read from a channel which no one is writing
//Assuming we do not write select statement to read from channel

package main

import (
	"fmt"
)

func main() {
	a := make(chan bool)
	result := <-a
	fmt.Println(result)
}

/*
  output -
  go run deadlock.go
  fatal error: all goroutines are asleep - deadlock!

  goroutine 1 [chan receive]:
  main.main()
          /home/yogesh/Go/src/goroutine_lim_patt/deadlock.go:11 +0x57
  exit status 2
*/
