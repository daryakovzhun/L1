package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type Counters struct {
	mx sync.Mutex  // поле для mutex
	m  map[int]int // мапа, где будут храниться значения
}

func NewCounters() *Counters {
	return &Counters{ //возвращаем ссылку на новый объект типа Counters
		m: make(map[int]int), // создаем мапу
	}
}

func (c *Counters) Store(key int, value int) {
	c.mx.Lock()         // блокируем для записи
	defer c.mx.Unlock() // отложенный вызов разблокировки
	c.m[key] = value    // добавляем пару ключ-значение
}

func main() {
	//  создаем массив с данными
	arr := []int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}

	//  1 способ - использование sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(1)                 //  добавляем 1 горутину
	counters := NewCounters() //  создаем объект структуры Counters, в которой хранится map
	go func() {               // запускаем горутину
		for k, v := range arr { //  итерируемся по массиву значений
			counters.Store(k, v) // конкуретно записываем в структуру
		}
		wg.Done() //  вычитаем из wg
	}()

	// 2 способ - использование sync.Map
	var counters1 sync.Map // создаем переменную типа sync.Map для конкуретной записи
	wg.Add(1)              //  добавляем 2 горутину
	go func() {            // запускаем горутину
		for k, v := range arr { //  итерируемся по массиву значений
			counters1.Store(k, v) // вызываем метод для конкуретной записи
		}
		wg.Done() //  вычитаем из wg
	}()

	wg.Wait() // дожидаемся завершения работ горутин

	fmt.Println("sync.Mutex")      //  выводим на экран название метода, который использовали
	for k, v := range counters.m { // итерируемся по map из структуры
		fmt.Println("Key = ", k, " , value = ", v) // выводим значения, которые были записаны
	}

	fmt.Println("\nsync.Map")                     //  выводим на экран название метода, который использовали
	counters1.Range(func(k, v interface{}) bool { // функция Range, принимающая анонимную функцию, вызываемую для каждого элемента мапы
		fmt.Println("key:", k, ", val:", v) // выводим значения, которые были записаны
		return true                         // если функция возвращает false, итерирование прекращается
	})
}
