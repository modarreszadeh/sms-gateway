package db

import (
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"github.com/modarreszadeh/sms-gateway/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func Seed(mongoDb *mongo.Database) {
	db = mongoDb

	seedUser()
	seedSmsConfig()
}

func seedUser() {
	userService := service.NewUserService(db)

	if userService.IsAnyUser() {
		return
	}

	var user = domain.NewUser("mohammad", "09393639116")

	userId, err := userService.CreateUser(user)
	if err != nil {
		return
	}

	seedUserSmsConfig(userId)
}

func seedUserSmsConfig(userId string) {
	userSmsConfigService := service.NewUserSmsConfigService(db)

	var userSmsConfig = domain.NewUserSmsConfig(userId, 0)
	err := userSmsConfigService.CreateUserSmsConfig(userSmsConfig)
	if err != nil {
		return
	}
}

func seedSmsConfig() {
	smsConfigService := service.NewSmsConfigService(db)

	if smsConfigService.IsAnySmsConfig() {
		return
	}

	var minValidChar byte = 4
	var maxValidChar = 250
	var charCost = 20
	smsConfig := domain.NewSmsConfig(minValidChar, maxValidChar, charCost)

	err := smsConfigService.CreateSmsConfig(smsConfig)
	if err != nil {
		return
	}
}
