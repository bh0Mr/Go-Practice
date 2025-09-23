package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("This runs at the end!")
	}()

	fmt.Println("Main logic running")

	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 4 is", square(4))

	multiply := func(a, b int) int {
		return a * b
	}
	result := multiply(4, 5)
	fmt.Println("4 * 5 =", result)


	fmt.Println("=====================================")

	Operation(3, 5, func(x, y int) int {
		return x + y
	})

	Operation(10, 2, func(x, y int) int {
		return x - y
	})
}






func Operation(a, b int, op func(int, int) int) {
	result := op(a, b)
	fmt.Println("Result:", result)
}
