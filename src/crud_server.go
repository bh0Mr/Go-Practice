package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	

)

type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Employees int   `json:"employees"`
}

var companies = []Company{}
var idCounter = 1

func GetCompanys(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method !=http.MethodGet {
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


func GetCompany(w http.ResponseWriter,r  * http.Request){
	w.Header().Set("Content-Type", "application/json")


	if r.Method !=http.MethodGet {
		w.Header().Set("X-Error-Message", "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == ""{
	   http.Error(w, "ID parameter is required", http.StatusBadRequest)
	   return
	}

	for _,c := range companies {
		if fmt.Sprintf("%d",c.ID) == id {
			json.NewEncoder(w).Encode(c)
			return

		}
	}

	http.Error(w,"Company not founf",http.StatusNotFound)


}
func AddCompany(w http.ResponseWriter, r *http.Request) {

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


	newCompany.ID = idCounter
	idCounter++
	companies = append(companies, newCompany)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompany)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        w.Header().Set("X-Error-Message", "Invalid request method")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return 
    }

    var updatedCompany Company
    err := json.NewDecoder(r.Body).Decode(&updatedCompany)
    if err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }

    for i, c := range companies {
        if c.ID == updatedCompany.ID {
            companies[i] = updatedCompany
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(updatedCompany)
            return
        }
    }

    http.Error(w, "Company not found", http.StatusNotFound)
}


func DeleteCompany(w http.ResponseWriter, r *http.Request) {

	if r.Method !=http.MethodDelete {
		w.Header().Set("X-Error-Message", "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }

    for i, c := range companies {
        if fmt.Sprintf("%d", c.ID) == id {

            companies = append(companies[:i], companies[i+1])
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(companies)

            return
        }
    }

    http.Error(w, "Company not found", http.StatusNotFound)
}

func main() {



	http.HandleFunc("/GetCompanys", GetCompanys)
	http.HandleFunc("/GetCompany", GetCompany)
	http.HandleFunc("/AddCompany", AddCompany)
	http.HandleFunc("/UpdateCompany", UpdateCompany)
	http.HandleFunc("/DeleteCompany", DeleteCompany)
	log.Println("Server started at http://localhost:8080")


    http.ListenAndServe(":8080", nil)
}
