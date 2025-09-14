package main

import (
	"encoding/json"
	"fmt"
	"strings"
)



func main() {
	
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}



	fmt.Println("JSON:", string(jsonData))



	
	jsonString := `{"name":"Bob","age":25}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Println("Go Struct:", p2)






	fmt.Println("==================================")


	jsonData2 := `{"name":"A","grade":[90,50,45]}`

	var student Student

	err = json.Unmarshal([]byte(jsonData2),&student)  

	if err != nil {
		fmt.Println("UnMarshling fail")
		return
	}

	fmt.Println(student)

	S := Student{Name:"B",Grade:[]int{40,50,99}}

	jsonData3,err := json.Marshal(S)

	if err != nil {
		fmt.Println("Marchlng errot")
		return
	}

	fmt.Println(string(jsonData3))




	fmt.Println("==============================")


CompanyJson := `{"companyName":"ABC","location":"Uganda", "department":{"dname":"CS","dempolyes":42} }`

var CompanyData Company
w.Header().Set("Content-Type", "application/json")

err = json.Unmarshal([]byte(CompanyJson),&CompanyData)

if err != nil {
	fmt.Println("Error while Unmarshalling Data")
	return 
}

fmt.Println(CompanyData)




fmt.Println("============================================")




jsonString2 := `{"Name":"A","Age":20,"Id":"243"}`
 


reader := strings.NewReader(jsonString2)

var jsonData1 interface{}


err = json.NewDecoder(reader).Decode(&jsonData1)


if err != nil{

	fmt.Println("Error")
}


fmt.Println(jsonData1)



	





}




type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`

}


type Student struct {

	Name string `json:"name"`
	Grade []int `json:"grade"`
}




//embedded structure 


type Department struct{

	Dname string   `json:"dname"`
	Dempolyes int   `json:"dempolyes"`


}



type Company struct {

	CompanyName string  `json:"companyName"`
	Location string      `json:"location"`
	Department Department   `json:"department"`

}



//////





