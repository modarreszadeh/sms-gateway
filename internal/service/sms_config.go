package service

import (
	"context"
	"github.com/modarreszadeh/sms-gateway/internal/db/collection"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SmsConfigService struct {
	db *mongo.Database
}

func NewSmsConfigService(db *mongo.Database) *SmsConfigService {
	return &SmsConfigService{
		db: db,
	}
}

func (service *SmsConfigService) GetSmsConfig() (*domain.SmsConfig, error) {
	var smsConfig domain.SmsConfig

	err := service.db.Collection(collection.SmsConfig).FindOne(context.Background(), bson.D{}).Decode(&smsConfig)
	if err != nil {
		return nil, err
	}

	return &smsConfig, nil
}

func (service *SmsConfigService) CreateSmsConfig(smsConfig *domain.SmsConfig) error {
	_, err := service.db.Collection(collection.SmsConfig).InsertOne(context.Background(), smsConfig)
	if err != nil {
		return err
	}
	return nil
}

func (service *SmsConfigService) IsAnySmsConfig() bool {
	isExists := false
	result := service.db.Collection(collection.SmsConfig).FindOne(context.Background(), bson.D{{}})
	if result.Err() == nil {
		isExists = true
	}
	return isExists
}
