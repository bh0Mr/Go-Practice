package main

import "fmt"



func main(){

	fmt.Println(Max(1,2,3,4,5,6,7))

    Data(1,2,3,4,5)

    Cars("BME","Thar","Maruti")


}


func Max(nums...int) int {
    if len(nums) == 0 {
        return 0 
    }
    maxVal := nums[0]
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }
    return maxVal
}



func Data(nums ...int){

    if len(nums) == 0 {

        return 
    }

    for _,n := range nums{

        n := n+1



        fmt.Println(n)

    
       }

}


func Cars(car...string){

    if len(car)==0{
        return
    }

    for _,cr := range car{

        fmt.Println(cr)
    }




}

 

