package main

import (
    "go.uber.org/zap"
    "net/http"
    "strconv"
   

    "github.com/gin-gonic/gin"
    //"go.uber.org/zap/zapcore"

)
/*
func main() {
    logger, _ := zap.NewProduction()

    logger.Info("User logged in",
        zap.Int("user_id", 42),
        zap.String("ip", "192.168.1.1"),
    )
} */

// User represents a simple user model
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}


var users = make(map[int]User)
var nextID = 1


func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {


        c.Next() 

        status := c.Writer.Status()
        clientIP := c.ClientIP()
        method := c.Request.Method
        path := c.Request.URL.Path

        logger.Info("HTTP request",
            zap.Int("status", status),
            zap.String("method", method),
            zap.String("path", path),
            zap.String("client_ip", clientIP),
        
        )
    }
}

func main() {

    config := zap.NewProductionConfig()

    logger, err := config.Build()
    if err != nil {
        panic(err)
    } 
    defer logger.Sync()

    
    r := gin.New()
    r.Use(ZapLogger(logger))
    r.Use(gin.Recovery())


    r.GET("/users", GetUsers)
    r.GET("/users/:id", GetUserByID)
    r.POST("/users", CreateUser)
    r.PUT("/users/:id", UpdateUser)
    r.DELETE("/users/:id", DeleteUser)

    logger.Info("Starting server", zap.String("port", "8080"))
    r.Run(":8080")
}


func GetUsers(c *gin.Context) {
    list := make([]User, 0, len(users))
    for _, u := range users {
        list = append(list, u)
    }
    c.IndentedJSON(http.StatusOK, list)
}

func GetUserByID(c *gin.Context) {

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
        return
    }
    user, exists := users[id]
    if !exists {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    u.ID = nextID
    nextID++
    users[u.ID] = u
    c.IndentedJSON(http.StatusCreated, u)
}

func UpdateUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
        return
    }
    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if _, exists := users[id]; !exists {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    u.ID = id
    users[id] = u
    c.IndentedJSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
        return
    }
    if _, exists := users[id]; !exists {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    delete(users, id)
    c.IndentedJSON(http.StatusNoContent, nil)
}



