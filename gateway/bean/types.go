package bean

// KYCParams KYCParams
type KYCParams struct {
	AppID   int64  `form:"app_id" json:"app_id"`     // 外部应用ID
	IDType  int8   `form:"id_type" json:"id_type"`   // 外部应用ID|Key 0:int64,1:string
	UserID  int64  `form:"user_id" json:"user_id"`   // 外部用户ID,int64
	UserKey string `form:"user_key" json:"user_key"` // 外部用户ID,string
	UID     int64  `form:"uid" json:"uid"`           // 用户ID
	Name    string `form:"name" json:"name"`         // 姓名
	RoleID  int8   `form:"role_id" json:"role_id"`   // 用户角色ID
	Pwd     string `form:"pwd" json:"pwd"`           // 用户交易密码
	OldPwd  string `form:"old_pwd" json:"old_pwd"`   // 用户旧交易密码
}

// TxParams TxParams
type TxParams struct {
	UID       int64  `form:"uid" json:"uid"`               // 用户ID
	Sender    int64  `form:"sender" json:"sender"`         // sender
	Receiver  int64  `form:"receiver" json:"receiver"`     // receiver
	TxTypeID  int8   `form:"tx_type_id" json:"tx_type_id"` // 交易类型ID
	Nonce     int64  `form:"nonce" json:"nonce"`           // sender nonce
	TokenID   int64  `form:"token_id" json:"token_id"`     // 币种ID
	Value     int64  `form:"value" json:"value"`           // 操作金额
	Pwd       string `form:"pwd" json:"pwd"`               // 用户交易密码
	ExAppID   int64  `form:"ex_app_id" json:"ex_app_id"`   // 外部应用ID
	PayloadID int64  `form:"payload_id" json:"payload_id"` // payload_id
	Payload   string `form:"payload" json:"payload"`       // payload
	Payment   string `form:"payment" json:"payment"`       // APP payment
	Page      int64  `form:"page" json:"page"`             // 查询page参数
	Memo      string `form:"memo" json:"memo"`             // 用户交易备注
}
