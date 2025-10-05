package main
import (

	"fmt"
	//"time"
	"sync"
)


func main(){

	/*go name()
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

	} */
	




    // Launching both the runners
    execute()

}

func runner1(wg *sync.WaitGroup) {
    fmt.Println("I am first runner")
	defer wg.Done()
}

func runner2(wg *sync.WaitGroup) {
    fmt.Println("I am second runner")
	defer wg.Done()
}


func execute() {
	wg := new(sync.WaitGroup)


	wg.Add(5)
    go runner1(wg)
    go runner2(wg)



	go task1(wg)
	go task2(wg)
	go task3(wg)

	wg.Wait()


	//time.Sleep(time.Second)


}


/////////////////////////////////////////////////////////////////////////

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 1 completed")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 2 completed")
	
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 3 completed")
}