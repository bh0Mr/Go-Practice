package main

import "fmt"

func example() {
    fmt.Println("Start of function")

    // This function will be called after the main function finishes.
    defer fmt.Println("This will be printed at the end")

    fmt.Println("Middle of function")
}

func main() {
    example()
}
