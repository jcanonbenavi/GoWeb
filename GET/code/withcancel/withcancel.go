package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

// asynchronous function that simulates an operation
func SimulatedOperation(ctx context.Context) {
	fmt.Println("Simulated operation")

	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing something...")
		//if context is canceled, the operation is canceled
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Operation finished")
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	go SimulatedOperation(ctx)
	//main is blocked until user press enter
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancelFunc()
	fmt.Println("End of the program")
}
