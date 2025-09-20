package main

import "fmt"




func main() {

    p := Person{
		Name: "Test",
	}
    
    fmt.Println(p.Speak())


	fmt.Println("=========================================")
 
	printType(42)
    printType("Hello Go")
    printType(3.14)

    fmt.Println("=========================================")

	doc := Document{
		Title: "Interface Practice",
	}
    PrintStart(doc)

	Img := Image{

		NumberOfImages:24,
	}
	PrintStart(Img)
	

    fmt.Println("=========================================")

	var worker Worker
    worker = Employee{
		Name: "Adam",
	}
    worker.Work()
    worker.Rest()
}


type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

////////////////////////////////////////////////////////////

func printType(value interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", value, value)
}

///////////////////////////////////////////////////////////

type Printer interface {
    Print()
}

type Document struct {
    Title string
}

type Image struct {
	NumberOfImages int 
}

func (d Document) Print() {
    fmt.Println("Document Title:", d.Title)
}

func (I Image) Print(){
	fmt.Println("Number of Iamges:",I.NumberOfImages)
}

func PrintStart(p Printer) {
    p.Print()
}

/////////////////////////////////////

type Worker interface {
    Work()
    Rest()
}


type Employee struct {
    Name string
}

func (e Employee) Work() {
    fmt.Println(e.Name + " is working")
}

func (e Employee) Rest() {
    fmt.Println(e.Name + " is resting")
}
