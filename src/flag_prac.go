package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "User's name")
	age := flag.Int("age", 18, "User's age")
	active := flag.Bool("active", false, "Is user active?")

	flag.Parse() 
	

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Active:", *active)
}
