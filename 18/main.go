package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в
//конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	i   int        // обычный счетчик
	mux sync.Mutex // mux для счетчика i

	i32 int32 // счетчик для atomic

	ichan   int      // счетчик с каналом
	chaninc chan int // канал счетчика
}

func NewCounter() *Counter {
	return &Counter{ //  создаем и возвращаем ссылку на новый объект типа Counter
		chaninc: make(chan int),
	}
}

func (si *Counter) AddMux() {
	si.mux.Lock()   // блокруем
	si.i++          // увеличиваем счетчик
	si.mux.Unlock() // разблокируем
}

func (si *Counter) AddAtomic() {
	atomic.AddInt32(&si.i32, 1) // с помощью atomic увеличиваем значение счетчика
}

func (si *Counter) readChan() {
	for range si.chaninc {
		si.ichan++ // инкрементируем счетчик пока канал открыт
	}
}

func (si *Counter) AddChan() {
	si.chaninc <- 1 // записываем в канал 1
}

func main() {
	wg := sync.WaitGroup{}     // для горутин в main
	chanWg := sync.WaitGroup{} // для канала в счетчике

	s := NewCounter()

	chanWg.Add(1) // добавляем chanWg 1 горутину

	go func() {
		s.readChan()  //  читаем канал
		chanWg.Done() //  вычитаем из chanWg
	}()

	wg.Add(2) // добавляем в wg 2 горутины

	go func() {
		for i := 0; i < 30; i++ {
			s.AddMux()    //  вызывем метод с импользованием mutex
			s.AddAtomic() //  вызывем метод с импользованием atomic
			s.AddChan()   //  вызывем метод с импользованием канала
		}
		wg.Done() // вычитаем из wg
	}()

	go func() {
		for i := 0; i < 20; i++ {
			s.AddMux()    //  вызывем метод с импользованием mutex
			s.AddAtomic() //  вызывем метод с импользованием atomic
			s.AddChan()   //  вызывем метод с импользованием канала
		}
		wg.Done() // вычитаем из wg
	}()

	wg.Wait() // дожидаемся завершения работы горутин

	close(s.chaninc) // закрываем канал
	chanWg.Wait()    // дожидаемся завершения работы горутины

	fmt.Println(s.i)     // выводим значение счетчика с импользованием mutex
	fmt.Println(s.i32)   // выводим значение счетчика с импользованием atomic
	fmt.Println(s.ichan) // выводим значение счетчика с импользованием канала
}
