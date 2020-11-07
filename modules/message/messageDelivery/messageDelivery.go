package messageDelivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"tes-warungpintar/models/jsonModel"
	"tes-warungpintar/models/messageModel"
	"tes-warungpintar/modules/message"
)

type messageHandler struct {
	messageUsecase message.MessageUsecase
}

//NewMessageHTTPHandler - will create new an object implementation of message.MessageUsecase interface
func NewMessageHTTPHandler(r *gin.Engine, messageUC message.MessageUsecase) {
	handler := messageHandler{
		messageUsecase: messageUC,
	}

	msg := r.Group("/api/v1/message")
	msg.Use()
	{
		msg.POST("/create", handler.SubmitMessage)
	}
}

// SubmitCreateAccount function for create account registration
// @Tags Account
// @Summary Create account registration
// @Description Create account registration
// @Accept json

// @Param body body messageModel.MessageRequest true "Body"
// @Produce json
// @Success 200 {object} jsonModel.JSONResponsePost{}
// @Failure 400 {object} jsonModel.JSONResponseBadRequest{}
// @Failure 401 {object} jsonModel.JSONErrorResponse{}
// @Failure 403 {object} jsonModel.JSONResponsePost{}
// @Failure 423 {object} jsonModel.JSONErrorResponse{}
// @Failure 500 {object} jsonModel.JSONErrorResponse{}
// @Router /message/create [post]
func (handler *messageHandler) SubmitMessage(c *gin.Context) {
	var myReq messageModel.MessageRequest
	err := c.BindJSON(&myReq)

	if err != nil {
		log.Error().Msg(err.Error())
		response := jsonModel.JSONErrorResponse{
			Status:   500,
			Messages: err.Error(),
			Target:   c.Request.RequestURI,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	errRes := handler.messageUsecase.AddMessage(&myReq)
	if errRes != nil {
		log.Error().Msg(errRes.Error())
		response := jsonModel.JSONErrorResponse{
			Status:   500,
			Messages: errRes.Error(),
			Target:   c.Request.RequestURI,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := jsonModel.JSONResponsePost{
		Status:   200,
		Messages: "Data has been saved",
	}

	c.JSON(http.StatusOK, response)
	return
}
