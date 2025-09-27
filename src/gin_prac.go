package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" `
}


var customers = []Customer{
	{ID: "1", Name: "test1", Email: "test1@example.com"},
	{ID: "2", Name: "test2", Email: "test2@example.com"},
}


func main() {
	router := gin.Default()

	route.Use(gin.Logger())
    route.Use(gin.Recovery())


	api := router.Group("/api")
	{
		api.GET("/customers", getCustomers)
		api.GET("/customers/:id", getCustomerByID)
		api.POST("/customers", createCustomer)
		api.PUT("/customers/:id", updateCustomer)
		api.DELETE("/customers/:id", deleteCustomer)
	}


	router.GET("/greet", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		c.String(http.StatusOK, "Hello, %s!", name)
	})


	router.GET("/context", func(c *gin.Context) {
		appName, _ := c.Get("appName")
		c.String(http.StatusOK, "This is %s", appName)
	})

	router.Run(":8080")
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("id")
	for _, cust := range customers {
		if cust.ID == id {
			c.IndentedJSON(http.StatusOK, cust)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func createCustomer(c *gin.Context) {
	var newCustomer Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

func updateCustomer(c *gin.Context) {
	id := c.Param("id")
	var updated Customer

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, cust := range customers {
		if cust.ID == id {
			customers[i] = updated
			c.IndentedJSON(http.StatusOK, updated)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")
	for i, cust := range customers {
		if cust.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Customer deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}
