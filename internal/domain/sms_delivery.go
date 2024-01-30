package domain

import (
	SmsStatus "github.com/modarreszadeh/sms-gateway/internal/domain/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SmsDelivery struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       string             `bson:"userId"`
	Sender       string             `bson:"sender"`
	Receptor     string             `bson:"receptor"`
	Message      string             `bson:"message"`
	Status       byte               `bson:"status"`
	Cost         int                `bson:"cost"`
	CreationDate time.Time          `bson:"creationDate"`
}

func NewSmsDelivery(userId string, sender string, receptor string, message string, cost int) *SmsDelivery {
	return &SmsDelivery{
		UserId:       userId,
		Sender:       sender,
		Receptor:     receptor,
		Message:      message,
		Cost:         cost,
		CreationDate: time.Now(),
		Status:       SmsStatus.InProgress,
	}
}
