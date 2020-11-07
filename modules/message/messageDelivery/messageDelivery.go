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
		msg.GET("/list/:userid", handler.GetListMessage)
	}
}

// SubmitCreateAccount function for create message
// @Tags Messages
// @Summary Create message
// @Description Create message
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

// GetListMessage function for get list message
// @Tags Messages
// @Summary Get List Message by user id
// @Description Get List Message
// @Accept json

// @Produce json
// @Success 200 {object} jsonModel.JSONResponse{Data=[]messageModel.MessageResponse{}}
// @Failure 400 {object} jsonModel.JSONResponseBadRequest{}
// @Failure 401 {object} jsonModel.JSONErrorResponse{}
// @Failure 403 {object} jsonModel.JSONResponsePost{}
// @Failure 423 {object} jsonModel.JSONErrorResponse{}
// @Failure 500 {object} jsonModel.JSONErrorResponse{}
// @Router /message/list/:userid [get]
func (handler *messageHandler) GetListMessage(c *gin.Context) {
	myUserID := c.Param("userid")
	if myUserID == "" {
		response := jsonModel.JSONResponsePost{
			Status:   400,
			Messages: "User id cant empty",
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	res, err := handler.messageUsecase.FindMessage(myUserID)
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

	if len(*res) == 0 {
		response := jsonModel.JSONResponse{
			Status:  200,
			Message: "Success",
			Data:    []messageModel.MessageResponse{},
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := jsonModel.JSONResponse{
		Status:  200,
		Message: "Success",
		Data:    &res,
	}

	c.JSON(http.StatusOK, response)
}
