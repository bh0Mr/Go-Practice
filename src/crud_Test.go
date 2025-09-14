package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct to represent a company
type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Employes int    `json:"emploes"`
}

// In-memory slice to hold company data
var companies []Company

// Handler for GET request (fetch all companies)
func getCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(companies) == 0 {
		w.WriteHeader(http.StatusOK) // Return 200 OK with empty array
		json.NewEncoder(w).Encode([]Company{})
		return
	}
	w.WriteHeader(http.StatusOK) // Return 200 OK
	json.NewEncoder(w).Encode(companies)
}

// Handler for POST request (add new company)
func addCompany(w http.ResponseWriter, r *http.Request) {
	var newCompany Company
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Return 400 Bad Request if invalid input
		return
	}

	companies = append(companies, newCompany)
	w.WriteHeader(http.StatusCreated) // Return 201 Created
	json.NewEncoder(w).Encode(newCompany)
}

// Handler for PUT request (update the first company)
func updateCompany(w http.ResponseWriter, r *http.Request) {
	if len(companies) == 0 {
		w.WriteHeader(http.StatusNotFound) // Return 404 Not Found if no companies exist
		return
	}

	var updatedCompany Company
	err := json.NewDecoder(r.Body).Decode(&updatedCompany)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Return 400 Bad Request if invalid input
		return
	}

	// Update the first company
	companies[0] = updatedCompany
	w.WriteHeader(http.StatusOK) // Return 200 OK
	json.NewEncoder(w).Encode(updatedCompany)
}

// Handler for DELETE request (delete the first company)
func deleteCompany(w http.ResponseWriter, r *http.Request) {
	if len(companies) == 0 {
		w.WriteHeader(http.StatusNotFound) // Return 404 Not Found if no companies exist
		return
	}

	// Remove the first company from the list
	companies = companies[1:]
	w.WriteHeader(http.StatusNoContent) // Return 204 No Content
}

func main() {
	http.HandleFunc("/companies", getCompanies)   // For GET
	http.HandleFunc("/submit", addCompany)        // For POST
	http.HandleFunc("/update", updateCompany)    // For PUT
	http.HandleFunc("/delete", deleteCompany)    // For DELETE

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
