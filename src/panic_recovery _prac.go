package main

import "fmt"

func main() {
    fmt.Println("Start")

	safe()
    fmt.Println("End") 


	fmt.Println("============================================")

	ArrayCheck()
	


}



func causePanic() {
    panic("Something went terribly wrong!")
}

func safe(){
	defer func(){

	r := recover()
	if r != nil {
		fmt.Println("Now you are Recoverd",r)
	}

}()

causePanic()
}







///////////////////////////////////////

func ArrayCheck(){
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from error:", r)
        }
    }()

    numbers := []int{1, 2, 3}
    fmt.Println(numbers[5])
}

//////////////////////////////////

