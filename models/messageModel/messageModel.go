package messageModel

import "github.com/jinzhu/gorm"

type (
	//MessageDatas - mapping table message_datas
	MessageDatas struct {
		gorm.Model
		UserID  int    `gorm:"column:user_id"`
		Message string `gorm:"column:message"`
	}

	//MessageRequest - struct for request message
	MessageRequest struct {
		UserID  int    `json:"user_id" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	//MessageResponse - struct for response message
	MessageResponse struct {
		Message string `json:"messages"`
	}
)
