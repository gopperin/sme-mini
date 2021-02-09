package mariadb

import (
	"github.com/jinzhu/gorm"
)

// GudpUser 用户表
type GudpUser struct {
	gorm.Model
	GudpUserBase
}

// GudpUserBase GudpUserBase
type GudpUserBase struct {
	UID       int64  `gorm:"unique;index:idx_uid;column:uid;" form:"uid" json:"uid"` // 用户ID
	NickName  string `gorm:"default:''" form:"nick_name" json:"nick_name"`           // 姓名
	Mobile    string `gorm:"default:''" form:"mobile" json:"mobile"`                 // 手机号
	Email     string `gorm:"default:''" form:"email" json:"email"`                   // email
	Pwd       string `gorm:"default:''" form:"pwd" json:"pwd"`                       // 密码
	SecretKey string `gorm:"default:''" form:"secret_key" json:"secret_key"`         // google secret key

	Stauts int64 `gorm:"default:1" form:"stauts" json:"stauts"` // 1:有效 2:无效

	ExNum  int64  `gorm:"default:0" form:"ex_num" json:"ex_num"`    // 扩展整形类型，自定义
	ExData string `gorm:"default:''" form:"ex_data" json:"ex_data"` // 扩展信息，自定义
	Memo   string `gorm:"default:''" form:"memo" json:"memo"`       // 备注
}

// GudpUserProfile 用户表
type GudpUserProfile struct {
	gorm.Model
	GudpUserProfileBase
}

// GudpUserProfileBase GudpUserProfileBase
type GudpUserProfileBase struct {
	UID       int64  `gorm:"unique;index:idx_uid;column:uid;" form:"uid" json:"uid"` // 用户ID
	Name      string `gorm:"default:''" form:"name" json:"name"`                     // 姓名
	IDCard    string `gorm:"column:idcard;default:''" form:"idcard" json:"idcard"`   // 身份证号
	BirthDay  string `gorm:"default:''" form:"birth_day" json:"birth_day"`           // 出生年月日
	Sex       string `gorm:"default:''" form:"sex" json:"sex"`                       // 性别
	Phone     string `gorm:"default:''" form:"phone" json:"phone"`                   // 手机号
	Address   string `gorm:"default:''" form:"address" json:"address"`               // 常住地址
	Telephone string `gorm:"default:''" form:"telephone" json:"telephone"`           // 固定电话

	IDCardFront string `gorm:"column:idcard_front;default:''" form:"idcard_front" json:"idcard_front"` // 身份证正面URI
	IDCardBack  string `gorm:"column:idcard_back;default:''" form:"idcard_back" json:"idcard_back"`    // 身份证背面URI
	LiveFace    string `gorm:"default:''" form:"live_face" json:"live_face"`                           // 活体照片

	KycState    int8   `gorm:"default:0" form:"kyc_state" json:"kyc_state"`          // kyc审核状态 0=未审核 1=提交审核 2=审核处理中 3=审核通过 4=审核驳回,可以重新提交 5=审核失败,不能提交
	KycAuditMsg string `gorm:"default:''" form:"kyc_audit_msg" json:"kyc_audit_msg"` // 审核结果说明

	ExNum  int64  `gorm:"default:0" form:"ex_num" json:"ex_num"`    // 扩展整形类型，自定义
	ExData string `gorm:"default:''" form:"ex_data" json:"ex_data"` // 扩展信息，自定义
	Memo   string `gorm:"default:''" form:"memo" json:"memo"`       // 备注
}
