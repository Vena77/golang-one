package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 1)

	go func() {
		for i := 1; i < 1001; i++ {
			ch1 <- "one"
			fmt.Println(<-ch1)
			ch1 <- "two"
			fmt.Println(<-ch1)
			if i == 5 {
				ch2 <- "Sigterm"
				fmt.Println(<-ch2)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	<-ch2
	fmt.Println("exiting")
}
