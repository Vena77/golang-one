package main

import (
	"fmt"
	"sync"
)

const count = 1000

var counter int
var mutex sync.Mutex

func main() {
	criticalSection()
	fmt.Println(counter)
}

func criticalSection() {
	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < count; i += 1 {
		go func() {
			counter += 1
		}()
	}
}
