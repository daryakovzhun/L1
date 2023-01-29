package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(22+32+42….) с использованием конкурентных вычислений.

func sumWithChan() {
	arr := []int{2, 4, 6, 8, 10}        // создаем массив с последовательностью чисел
	chanOut := make(chan int, len(arr)) // создаем канал
	wg := sync.WaitGroup{}

	for _, val := range arr { // итерируемся по массиву
		wg.Add(1)          // добавляем горутину
		go func(val int) { // вызываем горутину
			chanOut <- val * val // записываем в канал квадрат значения
			wg.Done()            // вычитаем из wg
		}(val)
	}

	wg.Wait()      // дожидаемся завершения работы всех горутин
	close(chanOut) // закрываем канал

	var sum int              // создаем переменную для суммы
	for v := range chanOut { // итерируемся по значениям в канале
		sum += v // добавляем к сумме
	}

	fmt.Println("Sum1 = ", sum) // выводим на экран сумму
}

func sumWithAtomic() {
	arr := []int{2, 4, 6, 8, 10} // создаем массив с последовательностью чисел
	wg := sync.WaitGroup{}

	var sum int32             // создаем переменную для суммы
	for _, val := range arr { // итерируемся по массиву
		wg.Add(1)          // добавляем горутину
		go func(val int) { // вызываем горутину
			atomic.AddInt32(&sum, int32(val*val)) // говорим atomic добавить к переменной sum val*val
			wg.Done()                             // вычитаем из wg
		}(val)
	}
	wg.Wait()                   // дожидаемся завершения работы всех горутин
	fmt.Println("Sum2 = ", sum) // выводим на экран сумму
}

func sumWithMutex() {
	arr := []int{2, 4, 6, 8, 10} // создаем массив с последовательностью чисел
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	var sum int               // создаем переменную для суммы
	for _, val := range arr { // итерируемся по массиву
		wg.Add(1)          // добавляем горутину
		go func(val int) { // вызываем горутину
			mu.Lock()        // блокируем
			sum += val * val // добавляем к sum val*val
			mu.Unlock()      // разблокируем
			wg.Done()        // вычитаем из wg
		}(val)
	}
	wg.Wait()                   // дожидаемся завершения работы всех горутин
	fmt.Println("Sum3 = ", sum) // выводим на экран сумму
}

func main() {
	sumWithChan()   // вызываем нахождение суммы с использованием каналов
	sumWithAtomic() // вызываем нахождение суммы с использованием atomic
	sumWithMutex()  // вызываем нахождение суммы с использованием mutex
}
