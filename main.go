package main

import (
    "github.com/gin-gonic/gin"
	"github.com/Interview-GPT/user-auth-service/controllers"
	"github.com/Interview-GPT/user-auth-service/initializers"

)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}