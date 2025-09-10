package main

import (
    "fmt"

    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("root handler called")
    fmt.Fprintf(w, "Hello, World!")
}


func loginHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("login handler called")
    fmt.Fprintf(w, "I am login !!\n")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("signup handler called")
    fmt.Fprintf(w, "I am signup !!\n")
}
func main() {
    http.HandleFunc("/", handler)

    http.HandleFunc("/login", loginHandler)

    http.HandleFunc("/signup", signupHandler)



    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
