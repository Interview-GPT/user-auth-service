package main

import (
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/Interview-GPT/user-auth-service/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main(){
	fmt.Println("Hello, world!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}