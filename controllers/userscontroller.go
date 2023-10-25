package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Interview-GPT/user-auth-service/initializers"
	"github.com/Interview-GPT/user-auth-service/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(c *gin.Context){
	//Get the email/pass off req body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{

		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
		})

		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//Create the user 
	user := models.User{Email: body.Email, Password: string(hash) }
	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{})
}


func Login(c *gin.Context){
	//Get the email and pass off req body

	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{

		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
		})

		return
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	//Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 *30).Unix(), 
	})

	tokenString, err := token.SignedString([]byte (os.Getenv("SECRET")))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Token",
		})

		return
	}

	//Send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}