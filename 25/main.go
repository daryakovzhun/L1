package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Hello")
	Sleep(5 * time.Second)
	fmt.Println("Bye")
}
