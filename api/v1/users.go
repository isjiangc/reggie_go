package v1

type UserLoginRequest struct {
	Phone string `json:"phone"` // 手机号
	Code  string `json:"code"`  // 验证码
}
type SendMsgRequest struct {
	Phone string `json:"phone"` // 手机号
}

type Users struct {
	Id       int64  `json:"id"`       // 主键
	Name     string `json:"name"`     // 姓名
	Phone    string `json:"phone"`    // 手机号
	Sex      string `json:"sex"`      // 性别
	IdNumber string `json:"idNumber"` // 身份证号
	Avatar   string `json:"avatar"`   // 头像
	Status   int    `json:"status"`   // 状态 0:禁用，1:正常
}
