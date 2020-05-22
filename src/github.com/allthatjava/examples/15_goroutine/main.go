package main

import (
	"fmt"
	"time"
)

func count(c int) {
	for i := 50; i < c; i++{
		fmt.Printf("goroutine... %d\n", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main(){

	// "go" is a thread managed by Go runtime
	go count(100)

	for i := 1; i < 50; i++{
		fmt.Printf("NOTgoroutine...... %d\n", i)
		time.Sleep(1000 * time.Millisecond)
	}
}
