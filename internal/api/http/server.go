package http

import (
	"github.com/labstack/echo/v4"
	"github.com/modarreszadeh/sms-gateway/internal/config"
	"github.com/modarreszadeh/sms-gateway/pkg/queue"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	db    *mongo.Database
	echo  *echo.Echo
	queue *queue.Queue
}

func New(db *mongo.Database, echo *echo.Echo, queue *queue.Queue) *Server {
	server := &Server{
		db:    db,
		echo:  echo,
		queue: queue,
	}

	v1 := server.echo.Group("/v1/api")

	sms := v1.Group("/sms", server.GetUserId)
	sms.GET("/delivery-list", server.GetUserSmsDeliveryHandler)
	sms.POST("/send", server.SendSmsHandler)

	user := v1.Group("/user", server.GetUserId)
	user.POST("/increase-balance", server.IncreaseUserBalance)

	return server
}

func (server *Server) Serve() {
	server.echo.Logger.Fatal(server.echo.Start(config.Port))
}
