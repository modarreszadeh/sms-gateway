package http

import (
	"github.com/labstack/echo/v4"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	"github.com/modarreszadeh/sms-gateway/internal/service"
	"net/http"
	"time"
)

func (server *Server) SendSmsHandler(c echo.Context) error {
	smsDeliveryService := service.NewSmsDeliveryService(server.db)
	smsConfigService := service.NewSmsConfigService(server.db)
	userSmsConfigService := service.NewUserSmsConfigService(server.db)

	smsDeliveryRequest := struct {
		Sender   string
		Receptor string
		Message  string
	}{}
	userId := getUserId(c)

	err := c.Bind(&smsDeliveryRequest)
	if err != nil {
		return InternalServerError(c, err)
	}

	smsConfig, _ := smsConfigService.GetSmsConfig()
	cost := smsConfig.CalculateMessageCost(smsDeliveryRequest.Message)

	if userSmsConfigService.HasInventory(userId, cost) == false {
		return JsonResult(c, http.StatusBadRequest, "Inventory not enough")
	}

	smsDelivery := domain.NewSmsDelivery(userId, smsDeliveryRequest.Sender,
		smsDeliveryRequest.Receptor, smsDeliveryRequest.Message, cost)

	deliveryId, err := smsDeliveryService.CreateSmsDelivery(smsDelivery)
	if err != nil {
		return InternalServerError(c, err)
	}

	server.queue.Enqueue(deliveryId)

	server.queue.DispatchProcess(server.SendSmsProcess)

	return Ok(c, "Sms send successfully")
}

func (server *Server) GetUserSmsDeliveryHandler(c echo.Context) error {
	return Ok(c, "delivery endpoint")
}

func (server *Server) IncreaseUserBalance(c echo.Context) error {
	userSmsConfigService := service.NewUserSmsConfigService(server.db)

	requestBody := struct {
		Balance int
	}{}
	userId := getUserId(c)

	err := c.Bind(&requestBody)
	if err != nil {
		return InternalServerError(c, err)
	}

	err = userSmsConfigService.IncreaseUserBalance(userId, requestBody.Balance)
	if err != nil {
		return InternalServerError(c, err)
	}

	return Ok(c, "Increase Balance Successfully")
}

func (server *Server) SendSmsProcess(process interface{}) {
	smsDeliveryService := service.NewSmsDeliveryService(server.db)
	userSmsConfigService := service.NewUserSmsConfigService(server.db)

	if deliveryId, ok := process.(string); ok {
		var smsDelivery, _ = smsDeliveryService.GetSmsDeliveryById(deliveryId)

		time.Sleep(3000 * time.Millisecond) // certain delay for send sms to operator (MCI-MTN)
		server.echo.Logger.Printf("sms with %s id and sender:[%s] send to specific Telecommunications Operator",
			smsDelivery.Id.Hex(), smsDelivery.Receptor)

		smsDeliveryService.ChangeSmsDeliveryStatus(deliveryId, domain.Delivered)

		userSmsConfig, err := userSmsConfigService.GetUserSmsConfigByUserId(smsDelivery.UserId)
		if err != nil {
			return
		}

		err = userSmsConfig.DecreaseBalance(smsDelivery.Cost)
		if err != nil {
			return
		}

		userSmsConfigService.UpdateUserSmsConfig(smsDelivery.UserId, userSmsConfig)
	}
}

func getUserId(c echo.Context) string {
	return c.Get("userId").(string)
}
