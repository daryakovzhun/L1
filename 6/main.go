package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//1 - завершение main функции и main горутины;
//2 - прослушивание всеми горутинами channel, при закрытии channel отправляется значение по
//умолчанию всем слушателям, при получении сигнала все горутины делают return;
//3 - завязать все горутины на переданный в них context.

func main() {
	// 1 - использование каналов
	message := make(chan string)

	go func() {
		time.Sleep(time.Second)
		message <- "Hello 1"
	}()
	fmt.Println(<-message)

	// 2 - использование context
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				fmt.Println("Work 2")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
