package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct must match the JSON sent by the server
type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Employes int    `json:"emploes"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer resp.Body.Close()

	var c Company
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		fmt.Println("Failed to decode JSON:", err)
		return
	}

	fmt.Println("Received data from local server:")
	fmt.Println("Name:", c.Name)
	fmt.Println("Location:", c.Location)
	fmt.Println("Employes:", c.Employes)
}
