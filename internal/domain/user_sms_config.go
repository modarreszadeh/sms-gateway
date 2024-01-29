package domain

import (
	"errors"
	"github.com/google/uuid"
)

type UserSmsConfig struct {
	UserId   string `bson:"userId"`
	ApiKey   string `bson:"apiKey"`
	IsActive bool   `bson:"isActive"`
	Balance  int    `bson:"balance"`
}

func NewUserSmsConfig(userId string, balance int) *UserSmsConfig {
	return &UserSmsConfig{
		UserId:   userId,
		Balance:  balance,
		ApiKey:   uuid.New().String(),
		IsActive: true,
	}
}

func (userSmsConfig *UserSmsConfig) IsValidApiKey(apiKey string) bool {
	return apiKey == userSmsConfig.ApiKey
}

func (userSmsConfig *UserSmsConfig) IncreaseBalance(value int) {
	userSmsConfig.Balance += value
}

func (userSmsConfig *UserSmsConfig) DecreaseBalance(value int) error {
	if value > userSmsConfig.Balance {
		return errors.New("value should be less than balance")
	}

	userSmsConfig.Balance -= value
	return nil
}
