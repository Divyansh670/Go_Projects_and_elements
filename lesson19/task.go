package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin")
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Users Route",
		})
	})

	r.GET("/user-profile", func(c *gin.Context) {
		u := User{
			ID:    "1",
			Name:  "Divyansh",
			Email: "divyansh@example.com",
		}
		c.JSON(http.StatusOK, u)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id)
	})

	r.Run(":8080")
}
