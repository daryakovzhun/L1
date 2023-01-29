package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

func main() {
	var workerCount int                         // переменная для количества воркеров
	fmt.Println("Please, enter workers count:") // просим ввести количество воркеров
	fmt.Scan(&workerCount)                      // считывает количество воркеров

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с функцией отмены

	chanWork := make(chan int) // создаем канал

	for i := 0; i < workerCount; i++ { // запускаем цикл по количеству воркеров
		go writer(ctx, chanWork)  // запускаем горутину для записи в канал
		go printer(ctx, chanWork) // запускаем горутину для чтения из канала
	}

	sigExit := make(chan os.Signal, 1)                    // создаем канал для сигнала от пользователя
	signal.Notify(sigExit, os.Interrupt, syscall.SIGTERM) // просим переслать в канал сигнал os.Interrupt
	<-sigExit                                             // ждем сигнала

	cancel() // вызываем ыункцию отмены контекста

	time.Sleep(time.Second) // засыпаем на секунду
	fmt.Println("Exit")     // выводим на экран о завершении программы
}

func printer(ctx context.Context, chanWork <-chan int) {
	for {
		select {
		case <-ctx.Done():
			break
		case val := <-chanWork:
			fmt.Println(val)
		}
	}
}

func writer(ctx context.Context, chanWork chan<- int) {
	for {
		i := rand.Intn(1000)
		select {
		case <-ctx.Done():
			break
		default:
			chanWork <- i
		}
	}
}
