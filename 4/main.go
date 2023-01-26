package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var worker_count int
	for {
		fmt.Println("Please, enter workers count:")
		_, err := fmt.Scan(&worker_count)
		if err != nil {
			fmt.Println("Try again")
			continue
		}
		break
	}

	ctx, cancel := context.WithCancel(context.Background())

	chan_work := make(chan int)
	for i := 0; i < worker_count; i++ {
		go printer(ctx, chan_work)
	}
	go writer(ctx, chan_work)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("Exit")
	close(chan_work)
}

func printer(ctx context.Context, chan_work <-chan int) {
	for {
		select {
		case <-ctx.Done():
			break
		case val := <-chan_work:
			fmt.Println(val)
		}
	}
}

func writer(ctx context.Context, chan_work chan<- int) {
	var i int
	for {
		select {
		case <-ctx.Done():
			break
		default:
			chan_work <- i
		}
		i++
	}
}
