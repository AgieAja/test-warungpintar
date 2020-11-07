package messageHttp

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	messageHandler "tes-warungpintar/modules/message/messageDelivery"
	messageRepo "tes-warungpintar/modules/message/messageRepository"
	messageUC "tes-warungpintar/modules/message/messageUsecase"
)

//MessageRoute - list route message
func MessageRoute(engine *gin.Engine, db *gorm.DB) {
	messageRepository := messageRepo.NewMessageRepository(db)
	messageUsecase := messageUC.NewMessageUsecase(messageRepository)
	messageHandler.NewMessageHTTPHandler(engine, messageUsecase)
}
