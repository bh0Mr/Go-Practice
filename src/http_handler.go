package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("This is server-side")  
	w.Header().Set("Content-Type", "text/html") 
	w.Header().Add("X-Custom-Header", "value1")
	w.Header().Add("X-Custom-Header", "value2")    
	fmt.Fprintln(w, "Hello, you are logged in") 
	
}

func main() {
	http.HandleFunc("/", loginHandler)

	log.Println("Server starting on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		log.Fatal("ListenAndServe: ", err)

	}
}
