package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

var books = []Book{
    {ID: "1", Title: "Test A", Author: "A", Price: 45.50},
    {ID: "2", Title: "Test B", Author: "B", Price: 37.99},
    {ID: "3", Title: "Test C", Author: "C", Price: 55.00},
}

func main() {
    r := gin.Default()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    v1 := r.Group("/api/v1")
    {
        v1.GET("/books", getBooks)
        v1.GET("/books/:id", getBookByID)
        v1.POST("/books", postBook)
        v1.DELETE("/book/:id",deleteBook)

    
    }

    r.Run(":8080")
}

func getBooks(c *gin.Context) {

    c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, b := range books {
        if b.ID == id {
            c.IndentedJSON(http.StatusOK, b)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBook(c *gin.Context) {
    var newBook Book

    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    books = append(books, newBook)
    
    c.IndentedJSON(http.StatusCreated, newBook)
}
func deleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, b := range books {
        if b.ID == id {
           
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
            return
        }
    }

    
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
