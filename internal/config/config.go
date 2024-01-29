package config

import "github.com/modarreszadeh/sms-gateway/pkg/mongodb"

var (
	Port  = ":5000"
	Mongo = &mongodb.Config{
		Host:     "localhost",
		Port:     "27017",
		UserName: "root",
		Password: "example",
		Database: "sms-gateway",
	}
)
