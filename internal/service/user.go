package service

import (
	"context"
	"github.com/modarreszadeh/sms-gateway/internal/db/collection"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserService struct {
	db *mongo.Database
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		db: db,
	}
}

func (service *UserService) GetUsers() []domain.User {
	var users []domain.User

	cursor, err := service.db.Collection(collection.User).Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	return users
}

func (service *UserService) CreateUser(user *domain.User) (userId string, err error) {
	result, err := service.db.Collection(collection.User).InsertOne(context.Background(), user)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Failed to convert inserted ID to ObjectID")
	}

	if err != nil {
		return insertedID.Hex(), err
	}
	return insertedID.Hex(), nil
}

func (service *UserService) CreateUserRange(users []interface{}) error {
	_, err := service.db.Collection(collection.User).InsertMany(context.Background(), users)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) IsAnyUser() bool {
	isExists := false
	result := service.db.Collection(collection.User).FindOne(context.Background(), bson.D{{}})
	if result.Err() == nil {
		isExists = true
	}
	return isExists
}
