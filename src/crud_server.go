package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Employees int   `json:"employees"`
}

var companies []Company

func GetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if len(companies) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Company{})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companies)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	var newCompany Company


	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}


	companies = append(companies, newCompany)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCompany)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	
	var updatedCompany Company

	err := json.NewDecoder(r.Body).Decode(&updatedCompany)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(companies) > 0 {
		companies[0] = updatedCompany
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(updatedCompany)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {


	companies = companies[1:]

	w.WriteHeader(http.StatusNoContent)
} 

func main() {

	http.HandleFunc("/Get", GetHandler)


	http.HandleFunc("/Post", PostHandler)

	http.HandleFunc("/Update", UpdateHandler)

	http.HandleFunc("/Delete", DeleteHandler)

	log.Println("Server started at http://localhost:8080")


    http.ListenAndServe(":8080", nil)
}
