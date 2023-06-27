package initializers

import "github.com/ashurmatovlochinbek/go-jwt/models"


func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}