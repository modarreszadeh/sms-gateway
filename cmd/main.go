package main

import (
	"github.com/labstack/echo/v4"
	"github.com/modarreszadeh/sms-gateway/internal/api/http"
	"github.com/modarreszadeh/sms-gateway/internal/config"
	"github.com/modarreszadeh/sms-gateway/internal/db"
	"github.com/modarreszadeh/sms-gateway/pkg/mongodb"
)

func main() {
	e := echo.New()

	mongoDatabase, _ := mongodb.NewMongoDbClient(config.Mongo)

	db.Seed(mongoDatabase)

	http.New(mongoDatabase, e).Serve()
}
