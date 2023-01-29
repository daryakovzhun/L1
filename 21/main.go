package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

// есть компьютер и существуют 2 вида mac и windows
type computer interface {
	insertInSquarePort()
}

// копьютер типа mac с методом insertInSquarePort
type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

// копьютер типа windows с методом insertInCirclePort
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

// у компьютера типа windows нет метода insertInSquarePort, хотя он выполняет такой же функционал как и метод
//insertInCirclePort, поэтому создаем адаптер для компьютера типа windows и реализуем недостающий метод (insertInSquarePort)
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

// есть клиент с компьютером типа windows или mac
type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

func main() {
	client := &client{}                   // создаем клиента
	mac := &mac{}                         // создаем компьютер типа mac
	client.insertSquareUsbInComputer(mac) // вызываем необходимый метод

	windowsMachine := &windows{}              // создаем компьютер типа windows
	windowsMachineAdapter := &windowsAdapter{ // создаем адаптер от компьютера типа windows
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter) // вызываем необходимый метод, которого у windows нет, но у windowsAdapter есть
}
