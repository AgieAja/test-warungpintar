package messageModel

import "github.com/jinzhu/gorm"

type (
	//MessageDatas - mapping table message_datas
	MessageDatas struct {
		gorm.Model
		UserID  string `gorm:"column:user_id"`
		Message string `gorm:"column:message"`
	}

	//MessageRequest - struct for request message
	MessageRequest struct {
		UserID  int    `json:"user_id"`
		Message string `json:"message"`
	}
)
