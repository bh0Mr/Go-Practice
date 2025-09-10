package main


import (
	"fmt"
	"encoding/json"
	"strings"
	"os"

) 


func main(){

	jsonString := `{"Name":"Test","Age":12,"ID":123,"MNo":12133}`


	reader := strings.NewReader(jsonString)

	var data Data

	decoder := json.NewDecoder(reader)

	//decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)

	if err != nil {

		fmt.Println("nil")

		return

	}


	fmt.Println(data)

	fmt.Println(decoder.InputOffset())




fmt.Println("========================================================")

//encoder 

Pdata := Person{
	Name:"A",
	Age:22,
	Id:111,
}

file ,_ := os.Create("output.json")

defer file.Close()



encoder := json.NewEncoder(file)

err = encoder.Encode(Pdata)

if err != nil {

	fmt.Println("Error")
	return 
}

fmt.Println(Pdata)






}


type Data struct {
	Name string
	Age  int 
	ID int 

}



////////encoder 


type Person struct {

	Name string `json:"name"`
	Age int  `json:"age`
	Id int `json:"id"`
}



