package main

import "fmt"

func main(){

	fmt.Println(Max(1,2,3,4,5,6,7))


}




func Max(nums ...int) int {
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
