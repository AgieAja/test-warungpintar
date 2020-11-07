package messageUsecase

import (
	"errors"
	"strconv"
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

//FindMessage - use case for find list message
func (msgUC *messageUsecase) FindMessage(userID string) (*[]messageModel.MessageResponse, error) {
	myUserID, errMyUserID := strconv.Atoi(userID)
	if errMyUserID != nil {
		return nil, errors.New("messageUC.FindMessage errMyUserID : " + errMyUserID.Error())
	}

	listData, errList := msgUC.messageRepository.RetrieveDatas(myUserID)
	if errList != nil {
		return nil, errList
	}

	if len(*listData) == 0 {
		return nil, nil
	}

	var list []messageModel.MessageResponse
	for _, val := range *listData {
		var detail messageModel.MessageResponse
		detail.Message = val.Message
		list = append(list, detail)
	}

	return &list, nil
}
