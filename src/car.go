package car

import "fmt"


type Car struct {
    Model string
    Year  int
    Owner string
}

func (c Car) DisplayDetails() {
    fmt.Println("Car Details:")
    fmt.Println("Model:", c.Model)
    fmt.Println("Year:", c.Year)
    fmt.Println("Owner:", c.Owner)
}


func (c *Car) UpdateOwner(newOwner string) {
    c.Owner = newOwner
}
