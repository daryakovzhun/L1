package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

type Point struct {
	_x, _y float64 //  координаты точки
}

func NewPoint(x, y float64) *Point {
	return &Point{ //  возвращаем ссылку на объект точки
		_x: x,
		_y: y,
	}
}

func (p *Point) GetX() float64 {
	return p._x //  возвращаем координату х точки
}

func (p *Point) GetY() float64 {
	return p._y //  возвращаем координату y точки
}

func DistanceBetweenPonts(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2)) //  расчитываем расстояние между точками
	// по формуле sqrt((x2 - x1)^2 + (y2 - y1)^2)
}

func main() {
	var x1, y1, x2, y2 float64     //  переменные для координат точек
	fmt.Print("Enter x1 and y1: ") //  просим ввести координаты 1 точки
	fmt.Scan(&x1, &y1)             //  считываем координаты 1 точки
	fmt.Print("Enter x2 and y2: ") //  просим ввести координаты 2 точки
	fmt.Scan(&x2, &y2)             //  считываем координаты 2 точки

	p1 := NewPoint(x1, y1)                      //  создаем 1 точку
	p2 := NewPoint(x2, y2)                      //  создаем 2 точку
	fmt.Println(DistanceBetweenPonts(*p1, *p2)) //  выводим посчитанное расстояние между точками
}
