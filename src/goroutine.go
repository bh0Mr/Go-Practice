package main
import (

	"fmt"
	"time"
)


func main(){

	go name()
	//time.Sleep(1 * time.Second)

	fmt.Println("This is main")



	go PrintNumber("A")
	go PrintNumber("B")

	//time.Sleep(1 * time.Second)


	fmt.Println("Finish Number Printing")


}

func name(){
	fmt.Println("This is go routine")
}






func PrintNumber(s string){

	for i := 0;i<3;i++ {
		fmt.Println(s,":",i)
		time.Sleep(1 * time.Second)

	}

}