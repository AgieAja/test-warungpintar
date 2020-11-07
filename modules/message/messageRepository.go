package message

import "tes-warungpintar/models/messageModel"

//MessageRepository - repository interface for module message
type MessageRepository interface {
	InsertData(data *messageModel.MessageRequest) error
}
