package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type SmsConfig struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	MinimumValidChar byte               `bson:"minimumValidChar"`
	MaximumValidChar int                `bson:"maximumValidChar"`
	CharacterCost    int                `bson:"characterCost"`
}

func NewSmsConfig(minValidChar byte, maxValidChar, charCost int) *SmsConfig {
	return &SmsConfig{
		MinimumValidChar: minValidChar,
		MaximumValidChar: maxValidChar,
		CharacterCost:    charCost,
	}
}

func (smsConfig *SmsConfig) IsValidCharacterMessage(message string) bool {
	if len(message) >= int(smsConfig.MinimumValidChar) && len(message) <= smsConfig.MaximumValidChar {
		return true
	}

	return false
}

func (smsConfig *SmsConfig) CalculateMessageCost(message string) int {

	return len(message) * smsConfig.CharacterCost
}
