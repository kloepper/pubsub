package main

import (
	"fmt"
	"time"

	"myelin"
)

type IntegerMessage int

func main() {
	names := []string{"A", "B", "C", "D"}
	for _, name := range names {
		messages := make(chan IntegerMessage)
		myelin.Subscribe(messages)
		printFormat := name + " got %d\n"
		go func() {
			for {
				fmt.Printf(printFormat, <-messages)
			}
		}()
	}

	for i := 0; i <= 25; i++ {
		myelin.Publish(IntegerMessage(i))
	}
	time.Sleep(time.Millisecond)
}
