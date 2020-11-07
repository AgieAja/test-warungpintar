package messageUsecase

import (
	"tes-warungpintar/models/messageModel"
	"tes-warungpintar/modules/message"
)

type messageUsecase struct {
	messageRepository message.MessageRepository
}

//NewMessageUsecase - will create new an messageUsecase object representation of message.MessageUsecase interface
func NewMessageUsecase(msgRepo message.MessageRepository) message.MessageUsecase {
	return &messageUsecase{
		messageRepository: msgRepo,
	}
}

//AddMessage - use case for add message
func (msgUC *messageUsecase) AddMessage(request *messageModel.MessageRequest) error {
	err := msgUC.messageRepository.InsertData(request)
	if err != nil {
		return err
	}

	return nil
}
