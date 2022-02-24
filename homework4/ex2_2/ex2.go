package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 1)

	// одна горутина читает значения из канала
	go func() {
		for val := range ch {
			fmt.Println(val)
			if val == "Sigterm" {
				break
			}
		}
	}()

	// другая горутина пишет в канал
	go func() {
		for i := 1; i < 40; i++ {
			ch <- "one"
			if i == 5 {
				ch <- "Sigterm"
				fmt.Println(<-ch)
				break
			}
		}
	}()

	// даём горутинам время выполниться
	time.Sleep(1 * time.Second)
}
