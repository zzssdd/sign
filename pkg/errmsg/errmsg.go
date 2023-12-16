package errmsg

const (
	SUCCESS = 200

	EmialExist           = 301
	EmailOrPasswordError = 302
	UserAlreadyInGroup   = 303

	ERROR = 500
)

var msg = map[int]string{
	SUCCESS: "成功",

	EmialExist:           "用户邮箱已被注册",
	EmailOrPasswordError: "邮箱或密码错误",
	UserAlreadyInGroup:   "用户已加入该群组",
	ERROR:                "失败",
}

func GetErrMsg(code int) string {
	return msg[code]
}
