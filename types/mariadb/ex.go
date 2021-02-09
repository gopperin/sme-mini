package mariadb

import (
	"github.com/jinzhu/gorm"
)

// GudpMessageTemplate GudpMessageTemplate
type GudpMessageTemplate struct {
	gorm.Model
	GudpMessageTemplateBase
}

// GudpMessageTemplateBase GudpMessageTemplateBase
type GudpMessageTemplateBase struct {
	Code   string `gorm:"default:''" form:"code" json:"code"`     // code sms:短信 email:邮件
	Title  string `gorm:"default:''" form:"title" json:"title"`   // title
	Detail string `gorm:"default:''" form:"detail" json:"detail"` // detail
	Lang   string `gorm:"default:''" form:"lang" json:"lang"`     // lang
	Type   string `gorm:"default:''" form:"type" json:"type"`     // type
	Memo   string `gorm:"default:''" form:"memo" json:"memo"`     // memo
}
