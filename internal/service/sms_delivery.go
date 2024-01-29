package service

import (
	"context"
	"github.com/modarreszadeh/sms-gateway/internal/db/collection"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"github.com/modarreszadeh/sms-gateway/pkg/mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SmsDeliveryService struct {
	db *mongo.Database
}

func NewSmsDeliveryService(db *mongo.Database) *SmsDeliveryService {
	return &SmsDeliveryService{
		db: db,
	}
}

func (service *SmsDeliveryService) GetAllUserSmsDelivery(userId string) []domain.SmsDelivery {
	var smsDeliveries []domain.SmsDelivery

	filter := bson.M{"userId": userId}
	cursor, err := service.db.Collection(collection.SmsDelivery).Find(context.Background(), filter)
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var smsDelivery domain.SmsDelivery
		if err := cursor.Decode(&smsDelivery); err != nil {
			return nil
		}
		smsDeliveries = append(smsDeliveries, smsDelivery)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	return smsDeliveries
}

func (service *SmsDeliveryService) GetSmsDeliveryById(smsDeliveryId string) (*domain.SmsDelivery, error) {
	var smsDelivery domain.SmsDelivery

	deliveryIdAsObjectId, _ := primitive.ObjectIDFromHex(smsDeliveryId)
	filter := bson.M{"_id": deliveryIdAsObjectId}

	err := service.db.Collection(collection.SmsDelivery).FindOne(context.Background(), filter).Decode(&smsDelivery)
	if err != nil {
		return nil, err
	}

	return &smsDelivery, nil
}

func (service *SmsDeliveryService) CreateSmsDelivery(smsDelivery *domain.SmsDelivery) (deliveryId string, err error) {
	result, err := service.db.Collection(collection.SmsDelivery).InsertOne(context.Background(), smsDelivery)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Failed to convert inserted ID to ObjectID")
	}

	if err != nil {
		return insertedID.Hex(), err
	}

	return insertedID.Hex(), nil
}

func (service *SmsDeliveryService) UpdateSmsDelivery(smsDeliveryId string, smsDelivery *domain.SmsDelivery) {
	deliveryIdAsObjectId, _ := primitive.ObjectIDFromHex(smsDeliveryId)
	filter := bson.M{"_id": deliveryIdAsObjectId}

	update := utils.GenerateBsonSetUpdate(map[string]interface{}{
		"sender":   smsDelivery.Sender,
		"receptor": smsDelivery.Receptor,
		"message":  smsDelivery.Message,
		"status":   smsDelivery.Status,
	})
	_, err := service.db.Collection(collection.SmsDelivery).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
}

func (service *SmsDeliveryService) ChangeSmsDeliveryStatus(smsDeliveryId string, status byte) {
	smsDelivery, _ := service.GetSmsDeliveryById(smsDeliveryId)

	smsDelivery.Status = status
	service.UpdateSmsDelivery(smsDeliveryId, smsDelivery)
}
