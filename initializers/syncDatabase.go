package initializers

import "github.com/Interview-GPT/user-auth-service/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}