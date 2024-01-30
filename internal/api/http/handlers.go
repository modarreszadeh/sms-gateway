package http

import (
	"github.com/labstack/echo/v4"
	"github.com/modarreszadeh/sms-gateway/internal/domain"
	SmsStatus "github.com/modarreszadeh/sms-gateway/internal/domain/enum"
	"github.com/modarreszadeh/sms-gateway/internal/service"
	"github.com/modarreszadeh/sms-gateway/pkg/queue"
	"net/http"
	"time"
)

var smsDeliveryQueue = queue.New(10)

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

	userSmsConfig, err := userSmsConfigService.GetUserSmsConfigByUserId(userId)
	if err != nil {
		return InternalServerError(c, err)
	}

	err = userSmsConfig.DecreaseBalance(cost)
	if err != nil {
		return InternalServerError(c, err)
	}

	userSmsConfigService.UpdateUserSmsConfig(userId, userSmsConfig)

	smsDelivery := domain.NewSmsDelivery(userId, smsDeliveryRequest.Sender,
		smsDeliveryRequest.Receptor, smsDeliveryRequest.Message, cost)

	deliveryId, err := smsDeliveryService.CreateSmsDelivery(smsDelivery)
	if err != nil {
		return InternalServerError(c, err)
	}

	smsDeliveryQueue.Enqueue(deliveryId)

	smsDeliveryQueue.DispatchProcess(server.SendSmsProcess)

	return Ok(c, "Sms send successfully")
}

func (server *Server) GetUserSmsDeliveryHandler(c echo.Context) error {
	smsDeliveryService := service.NewSmsDeliveryService(server.db)
	userId := getUserId(c)

	smsDeliveryList := smsDeliveryService.GetAllUserSmsDelivery(userId)

	return c.JSON(http.StatusOK, smsDeliveryList)
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

	if deliveryId, ok := process.(string); ok {
		var smsDelivery, _ = smsDeliveryService.GetSmsDeliveryById(deliveryId)

		time.Sleep(5000 * time.Millisecond) // certain delay for send sms to operator (MCI-MTN)
		server.echo.Logger.Printf("sms with %s id and sender:[%s] send to specific Telecommunications Operator",
			smsDelivery.Id.Hex(), smsDelivery.Receptor)

		smsDeliveryService.ChangeSmsDeliveryStatus(deliveryId, SmsStatus.Delivered)
	}
}

func getUserId(c echo.Context) string {
	return c.Get("userId").(string)
}
