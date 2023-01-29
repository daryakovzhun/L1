package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

func main() {
	var sec int                                             // переменная для времени выполнения программы
	fmt.Println("Please, enter time work script (second):") // просим ввести время в секундах
	fmt.Scan(&sec)                                          // считывем время

	wg := sync.WaitGroup{}
	wg.Add(2) // добавляем 2 горутины

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second) // создаем контекст, завязанный на времени sec

	chanWork := make(chan int) // создаем канал

	go func() { // запускаем горутину
		writer(ctx, chanWork) // вызываем функцию для отправки данных в канал
		wg.Done()             // вычитаем из wg
		close(chanWork)       // закрываем канал
	}()

	go func() { // запускаем горутину
		printer(ctx, chanWork) // вызываем функцию для печати данных из канала
		wg.Done()              // закрываем канал
	}()

	wg.Wait()           // дожидаемся завершения работы всех горутин
	fmt.Println("Exit") // выводим на экран о завершении программы
}

func printer(ctx context.Context, chanWork <-chan int) {
	for { //запускаем бесконечный цикл
		select { // оператор, который блокируется до тех пор, пока один из его блоков case не будет готов к запуску, а затем выполняет этот блок
		case <-ctx.Done(): // проверяем истекло ли время у контекста
			return // выходим из функции
		case val := <-chanWork: // записываем данные из канала в переменную
			fmt.Println(val) // выводим на экран полученные данные
		}
	}
}

func writer(ctx context.Context, chanWork chan<- int) {
	for { //запускаем бесконечный цикл
		i := rand.Intn(1000) // генинируем рандомное число до 1000
		select {             // оператор, который блокируется до тех пор, пока один из его блоков case не будет готов к запуску, а затем выполняет этот блок
		case <-ctx.Done(): // проверяем истекло ли время у контекста
			return // выходим из функции
		default:
			chanWork <- i // записываем данные в канал
		}
	}
}
