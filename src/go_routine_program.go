package main

import (
	"fmt"
	"sync"
	//"time" 
)

/*
var Count = 1



func Test1(){
	fmt.Println("Count1:",Count)
	Count++
}

func Test2(){
	fmt.Println("Count2:",Count)
	Count++
}


func Test3(){
	fmt.Println("Count3:",Count)
	Count++
}

func Test4(){
	fmt.Println("Count4:",Count)
	Count++
}

func Test5(){
	fmt.Println("Count5:",Count)
	Count++
}

func main(){


	var wg sync.WaitGroup

	wg.Add(5)
	defer wg.Done()

	go Test1()
	go Test2()
	go Test3()
	go Test4()
	go Test5()

	wg.Wait()



	fmt.Println("Code Finish")




}

	*/
	var (
		counter int
		mu      sync.Mutex
	)
	
	func increment(wg *sync.WaitGroup) {
		mu.Lock()
		counter++
		mu.Unlock()
		wg.Done()
	}
	
	func main() {
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go increment(&wg)
		}
		wg.Wait()
		fmt.Println("Counter:", counter)
	}





