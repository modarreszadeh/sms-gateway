package utils

import "go.mongodb.org/mongo-driver/bson"

func GenerateBsonSetUpdate(updateFields map[string]interface{}) bson.M {
	setUpdate := bson.M{"$set": bson.M{}}

	for key, value := range updateFields {
		setUpdate["$set"].(bson.M)[key] = value
	}

	return setUpdate
}
