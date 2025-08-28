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
class Bed {
    int length;
    int type;
    int breadth;
    int height;



}

type Door struct {

    length int
    height int 
    breadth int
    handel string
}
func (c *Car) UpdateOwner(newOwner string) {
    c.Owner = newOwner
}
