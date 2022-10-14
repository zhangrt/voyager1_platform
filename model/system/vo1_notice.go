package system

import "time"

// 配置文件结构体
type Vo1Notice struct {
	ID         string    `gorm:"primarykey" json:"id,string" form:"id"`
	Title      string    `json:"title" gorm:"comment:标题"`
	Module     string    `json:"module" gorm:"comment:消息体"`
	SenderId   string    `json:"sender_id" gorm:"comment:发送人ID"`
	Sender     string    `json:"sender" gorm:"comment:发送人"`
	NoticeTime time.Time `json:"notice_time" gorm:"comment:提醒时间"`
	Url        string    `json:"url" gorm:"comment:路径"`
	ReceiverId string    `json:"receiver_id" gorm:"comment:接收人ID"`
	Receiver   string    `json:"receiver" gorm:"comment:接收人"`
	UnRead     bool      `json:"unread" gorm:"comment:未读"`
	ReadTime   time.Time `json:"read_time" gorm:"comment:读取时间"`
}

func (Vo1Notice) TableName() string {
	return "vo1_notice"
}
