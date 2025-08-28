package main

import "fmt"

func example() {
    fmt.Println("Start of function")


    defer fmt.Println("This will be printed at the end")

    fmt.Println("Middle of function")
}

func main() {
    example()
}
