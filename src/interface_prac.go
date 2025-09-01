package main

import (
    "fmt"
    "math"
)

func main() {
    circ := Circle{Radius: 5}
    rect := Rectangle{Width: 4, Height: 3}

    PrintArea(circ) 
    PrintArea(rect) 

	//Animla example 

	d := Dog{}
    c := Cat{}

    MakeSound(d)
    MakeSound(c)

	// Vehicle example 

	C := Car{}
	B := Bike{}



	PrintWheel(C)
	PrintWheel(B)
	



}


type Shape interface {
    Area() float64
}


type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}


type Rectangle struct {
    Width, Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}


func PrintArea(s Shape) {
    fmt.Println("Area is:", s.Area())
}


//Animal sound 2nd example

type Animal interface {
    Sound() string
}

type Dog struct{}
func (Dog) Sound() string { 
	return "Woof" 
}

type Cat struct{}
func (Cat) Sound() string {
	return "Meow" 
}

func MakeSound(a Animal) {


    fmt.Println(a.Sound())
}



// Vehicle 3rd example 

type Vehicle interface {
	Wheel() string 
}

type Car struct {

}

func (Car) Wheel() string {

	return "Car has Four wheel"
}

type Bike struct {

}

func (Bike) Wheel() string{

	return "Bike has Two wheels "
}

func PrintWheel(V Vehicle){

	fmt.Println(V.Wheel())
}




