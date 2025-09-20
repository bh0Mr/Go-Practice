package main

import (
    "errors"
    "fmt"
)



func PracError(){

var erro = errors.New("SomeThing Fail")

wrraper1 := fmt.Errorf("This is also an error:%w",erro)
fmt.Println(wrraper1) 

wrraper2 := fmt.Errorf("This is 2nd error :%w",wrraper1)
fmt.Println(wrraper2)


unwrappedErr := errors.Unwrap(wrraper2)
	if unwrappedErr != nil {
		fmt.Println("Unwrapped error:", unwrappedErr)
	}

    if errors.Is(wrraper2, erro) {
		fmt.Println("The original error 'SomeThing Fail' is found in the error chain.")
	}



}


func divide(a, b int) (int, error) {
    if b == 0 {
                   return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil 
}


/////////////////////////////////////////////////////
type MyError struct {
    Code    int
    Message string
}
 



func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func mayFail() error {
    return &MyError{
        Code:    404,
        Message: "Resource not found",
    }
}




func PracError2(){
    err := mayFail()
    if err != nil {
        fmt.Println("Custom Error:", err)
    }

}












func main() {
    result, err := divide(10, 0) 

    if err != nil {
        fmt.Println("Error:", err) 
    } else {
        fmt.Println("Result:", result) 
    }


fmt.Println("================================================")



//fmt.Println(erro) 
PracError()

fmt.Println("================================================")

PracError2()




}



