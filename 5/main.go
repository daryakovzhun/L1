package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

func main() {
	var sec int
	for {
		fmt.Println("Please, enter time work script (second):")
		_, err := fmt.Scan(&sec)
		if err != nil {
			fmt.Println("Try again")
			continue
		}
		break
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)

	chan_work := make(chan int)

	go func() {
		writer(ctx, chan_work)
		wg.Done()
		close(chan_work)
	}()

	go func() {
		printer(ctx, chan_work)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Exit")
}

func printer(ctx context.Context, chan_work <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
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
			return
		default:
			chan_work <- i
		}
		i++
	}
}
