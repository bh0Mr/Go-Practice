package main

import "fmt"



func main() {
    
    person := &Person{Name: "Alice", Age: 30}



    fmt.Println("Before Birthday:", person.Age) 

    person.Birthday() 

    fmt.Println("After Birthday:", person.Age)   
}

type Person struct {
	 Name string
	 Age int
	  } 
	  
 func (p *Person) Birthday() { 
	 p.Age += 1 
	
	}


