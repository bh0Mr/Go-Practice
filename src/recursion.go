package main

import "fmt"

func main() {
	fmt.Printf("Factorial of number is: %d\n", factorial(5))
}



func factorial(n int) int {
	
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
r