package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

// Define the Company struct
type Company struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Employees int    `json:"employees"`
}

// Define a map to store companies
var companies = make(map[int]Company)
var companyIDCounter = 1

// Function to handle fetching a company by ID
func getCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the 'id' query parameter from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	// Search for the company with the given ID
	for _, company := range companies {
		if fmt.Sprintf("%d", company.ID) == id {
			// Company found, return it as JSON
			json.NewEncoder(w).Encode(company)
			return
		}
	}

	// If no company is found, return a 404
	http.Error(w, "Company not found", http.StatusNotFound)
}

// Get all companies (unchanged)
func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.Header().Set("X-Error-Message", "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(companies) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Company{})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

// Add a new company (unchanged)
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("X-Error-Message", "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newCompany Company
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	companies[companyIDCounter] = newCompany
	companyIDCounter++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompany)
}

// Update a company (unchanged)
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("X-Error-Message", "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCompany Company
	err := json.NewDecoder(r.Body).Decode(&updatedCompany)
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// For simplicity, assuming we update the company with ID = 1
	if len(companies) > 0 {
		companies[0] = updatedCompany
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCompany)
}

// Delete a company (unchanged)
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// This example deletes the first company (just for illustration)
	if len(companies) > 0 {
		for id := range companies {
			delete(companies, id)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Main function
func main() {
	// Route for getting all companies
	http.HandleFunc("/Get", GetCompany)

	// Route for creating a new company
	http.HandleFunc("/Post", PostHandler)

	// Route for updating a company
	http.HandleFunc("/Update", UpdateHandler)

	// Route for deleting a company
	http.HandleFunc("/Delete", DeleteHandler)

	// New route for getting a company by ID
	http.HandleFunc("/GetCompanyByID", getCompanyByID)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
