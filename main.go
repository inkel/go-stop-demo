package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ticker := time.NewTicker(1 * time.Second)

	ch := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()

	go func() {
		time.Sleep(10 * time.Second)
		stop()
	}()

	var i int

	fn := func() error {
		for {
			select {
			case <-ch:
				fmt.Println("STOPPED")
				return nil
			case <-ticker.C:
				i++
				fmt.Println("TICK", i)
			case <-ctx.Done():
				fmt.Println("CONTEXT")
				return ctx.Err()
			}
		}
	}

	fmt.Println("BYE", fn())
}
