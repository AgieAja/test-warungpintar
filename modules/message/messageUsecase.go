package message

import "tes-warungpintar/models/messageModel"

type MessageUsecase interface {
	AddMessage(request *messageModel.MessageRequest) error
}
