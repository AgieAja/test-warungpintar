package messageRepository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"tes-warungpintar/models/messageModel"
	"tes-warungpintar/modules/message"
)

type sqlRepository struct {
	Conn *gorm.DB
}

//NewMessageRepository - will create an object that representation that message.MessageRepository
func NewMessageRepository(Conn *gorm.DB) message.MessageRepository {
	return &sqlRepository{Conn}
}

//InsertData - insert data to table message_datas
func (db *sqlRepository) InsertData(data *messageModel.MessageRequest) error {
	messageData := messageModel.MessageDatas{
		UserID:  data.UserID,
		Message: data.Message,
	}

	err := db.Conn.Create(&messageData).Error
	if err != nil {
		return errors.New("messageRepo.InsertData err : " + err.Error())
	}

	return nil
}

//RetrieveDatas - get data from table message_datas
func (db *sqlRepository) RetrieveDatas(userID int) (*[]messageModel.MessageDatas, error) {
	var data []messageModel.MessageDatas
	err := db.Conn.Select("message").Where("user_id = ?", userID).Find(&data).Error
	if err != nil {
		return nil, errors.New("messageRepo.RetrieveDatas err : " + err.Error())
	}

	return &data, nil
}
