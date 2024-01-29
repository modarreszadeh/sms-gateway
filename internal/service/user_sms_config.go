package service

import (
	"context"
	"github.com/modarreszadeh/sms-gateway/internal/db/collection"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"github.com/modarreszadeh/sms-gateway/pkg/mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSmsConfigService struct {
	db *mongo.Database
}

func NewUserSmsConfigService(db *mongo.Database) *UserSmsConfigService {
	return &UserSmsConfigService{
		db: db,
	}
}

func (service *UserSmsConfigService) GetUserSmsConfigByUserId(userId string) (*domain.UserSmsConfig, error) {
	var userSmsConfig domain.UserSmsConfig

	filter := bson.M{"userId": userId}
	err := service.db.Collection(collection.UserSmsConfig).FindOne(context.Background(), filter).Decode(&userSmsConfig)
	if err != nil {
		return nil, err
	}

	return &userSmsConfig, nil
}

func (service *UserSmsConfigService) CreateUserSmsConfig(userSmsConfig *domain.UserSmsConfig) error {
	_, err := service.db.Collection(collection.UserSmsConfig).InsertOne(context.Background(), userSmsConfig)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserSmsConfigService) UpdateUserSmsConfig(userId string, userSmsConfig *domain.UserSmsConfig) {
	filter := bson.M{"userId": userId}

	update := utils.GenerateBsonSetUpdate(map[string]interface{}{
		"userId":   userId,
		"apiKey":   userSmsConfig.ApiKey,
		"balance":  userSmsConfig.Balance,
		"isActive": userSmsConfig.IsActive,
	})
	_, err := service.db.Collection(collection.UserSmsConfig).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
}

func (service *UserSmsConfigService) IncreaseUserBalance(userId string, value int) error {
	userSmsConfig, err := service.GetUserSmsConfigByUserId(userId)
	if err != nil {
		return err
	}

	userSmsConfig.IncreaseBalance(value)
	service.UpdateUserSmsConfig(userId, userSmsConfig)

	return nil
}

func (service *UserSmsConfigService) DecreaseUserBalance(userId string, value int) error {
	userSmsConfig, err := service.GetUserSmsConfigByUserId(userId)
	if err != nil {
		return err
	}

	err = userSmsConfig.DecreaseBalance(value)
	if err != nil {
		return err
	}

	service.UpdateUserSmsConfig(userId, userSmsConfig)

	return nil
}

func (service *UserSmsConfigService) HasInventory(userId string, cost int) bool {
	hasInventory := false

	userSmsConfig, _ := service.GetUserSmsConfigByUserId(userId)
	if userSmsConfig.Balance >= cost {
		hasInventory = true
	}

	return hasInventory

}
