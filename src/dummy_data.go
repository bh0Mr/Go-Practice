package main

import (

	"fmt"
	"encoding/json"
	"log"
	"net/http"
)



type Company struct {

	Name string `json:"name"`
	Location string `json:"location"`
	Employes int `json:"emploes"`

}

func loginHandler(w http.ResponseWriter, r *http.Request){
log.Println("This is server side ")

w.Header().Set("Content-Type", "application/json")


data := Company{

	Name:"ABC",
	Location:"Uganda",
	Employes: 1233,

}


err := json.NewEncoder(w).Encode(data)

if err != nil{
	fmt.Println("Error")
	return
}

}
func main(){

	http.HandleFunc("/",loginHandler)

	log.Println("Server start at port 8080")

	err := http.ListenAndServe(":8080",nil)

	if err != nil{

		log.Println("error")


	}


}