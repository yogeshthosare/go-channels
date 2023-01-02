//________________________________
//What will happen if we write in a channel which no one reading

package main

import (
	"fmt"
)

func main() {
	a := make(chan bool)
	a <- true
	fmt.Println("hay")

}

/*
  output -
  go run deadlock.go
  fatal error: all goroutines are asleep - deadlock!

  goroutine 1 [chan send]:
  main.main()
          /home/yogesh/Go/src/goroutine_lim_patt/deadlock.go:9 +0x54
  exit status 2
*/
