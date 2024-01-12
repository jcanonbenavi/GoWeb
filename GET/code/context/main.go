package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "my-value")
	ctx = context.WithValue(ctx, "my-key2", 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go myProcess(ctx2)

	<-ctx2.Done()
	fmt.Println("Context is done")
	fmt.Println(ctx2.Err())
}

func viewContext(ctx context.Context) {
	fmt.Printf("My value is %s\n", ctx.Value("my-key"))
	fmt.Printf("My value is %d\n", ctx.Value("my-key2"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done")
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
