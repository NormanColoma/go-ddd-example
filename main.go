package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func handle(c *gin.Context) {
	user := User{
		Name: "pepe",
	}
	c.JSON(http.StatusOK, user)
}

func main() {
	r := gin.Default()
	r.GET("/ping", handle)
	r.Run()
}
