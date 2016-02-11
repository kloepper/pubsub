package main

import (
	"fmt"
	"time"

	"myelin"
)

type foo struct {
	N int
}

func main() {
	names := []string{"A", "B", "C", "D"}
	for _, name := range names {
		messages := make(chan foo)
		myelin.Subscribe(messages)
		printFormat := name + " got foo.N=%d\n"
		go func() {
			for {
				fmt.Printf(printFormat, (<-messages).N)
			}
		}()
	}

	for i := 0; i <= 25; i++ {
		myelin.Publish(foo{N: i})
	}

	time.Sleep(time.Millisecond)
}
