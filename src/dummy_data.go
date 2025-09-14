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

func RootHandler(w http.ResponseWriter, r *http.Request){
log.Println("This is server side ")

w.Header().Set("Content-Type", "application/json")



 if r.Method  == http.MethodGet {

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
}

func PostHandler(w http.ResponseWriter,r *http.Request) {
	log.Println("This Server side Post Request Handler ")

	w.Header().Set("Content-Type","application/json")


	if r.Method == http.MethodPost {

		var receive Company

	err := json.NewDecoder(r.Body).Decode(&receive)

	if err != nil {

		log.Println("error")

		return
	}

	log.Printf("Received POST Data: Name=%s, Location=%s, Employes=%d",
			receive.Name, receive.Location, receive.Employes)



	} 



}



func main(){

	http.HandleFunc("/",RootHandler)

	http.HandleFunc("/Submit",PostHandler)

	log.Println("Server start at port 8080")

	err := http.ListenAndServe(":8080",nil)

	if err != nil{

		log.Println("error")


	}
	

	


}