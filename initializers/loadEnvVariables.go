package initializers

import (
    "log"
    "github.com/joho/godotenv"
)

//set up your own .env variables or ask owner for variables necessary for this service to work

func LoadEnvVariables(){
    
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

}