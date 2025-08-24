package main

import (
    "fmt"
    "/home/ubuntu/Documents/experiments/Go-Practice/car.go"  
)

func main() {
    /
    car1 := car.Car{,
        Model: "BMW",
        Year:  2020,
        Owner: "Rin",
    }

  
    car1.DisplayDetails()


    car1.UpdateOwner("Rin")


    fmt.Println("\nUpdated Car Details:")
    car1.DisplayDetails()
}
