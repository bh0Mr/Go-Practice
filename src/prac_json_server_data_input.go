package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Struct must match the JSON sent by the server
type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Employes int    `json:"emploes"`
}


type Detail struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Employes int    `json:"emploes"`
}






func TestPost(){

	data := Detail{
		Name:     "XYZ",
		Location: "India",
		Employes: 456,
	}

	jsonData,err := json.Marshal(data)
	if err != nil {

		fmt.Println("error while Marshal")
		return
	}

	

	resp, err := http.Post("http://localhost:8080/Submit","application/json",strings.NewReader(string(jsonData)))
	if err != nil {

		fmt.Println("error while Post Request")
		return
	}

	defer resp.Body.Close()

	fmt.Println("POST Response Status:", resp.Status)





}

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("Failed to connect:")
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	var c Company


	err = json.NewDecoder(resp.Body).Decode(&c)



	if err != nil {


		fmt.Println("Failed to decode JSON:")
		return
	}

	fmt.Println("Received data from local server:")

	fmt.Println("Name:", c.Name)

	fmt.Println("Location:", c.Location)
	
	fmt.Println("Employes:", c.Employes)

	TestPost()



	//////


}
