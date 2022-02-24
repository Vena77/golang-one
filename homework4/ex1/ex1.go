package main

import "fmt"

func main() {

	for i := 1; i < 1001; i++ {

		go func(n int) {
			fmt.Println(n, "-", n)
		}(i)
	}
	fmt.Scanln()
	fmt.Println("The End")
}
