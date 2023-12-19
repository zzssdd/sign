package errmsg

const (
	SUCCESS = 200
	Thanks  = 201

	EmialExist           = 301
	EmailOrPasswordError = 302
	UserAlreadyInGroup   = 303
	OrderExpired         = 304
	OrderFinished        = 305
	NotInTime            = 306
	ScoreNotEniough      = 307
	UserNotInGroup       = 308
	TokenNotExist        = 309
	TokenFormatError     = 310
	TokenInValid         = 311

	ERROR = 500
)

var msg = map[int]string{
	SUCCESS: "成功",
	Thanks:  "很遗憾，您未能中将",

	EmialExist:           "用户邮箱已被注册",
	EmailOrPasswordError: "邮箱或密码错误",
	UserAlreadyInGroup:   "用户已加入该群组",
	OrderExpired:         "订单不存在或已过期",
	OrderFinished:        "请勿重复提交",
	NotInTime:            "不在抽奖时间范围内",
	ScoreNotEniough:      "积分不足",
	UserNotInGroup:       "请先加入该群组",
	TokenNotExist:        "请先登陆",
	TokenFormatError:     "token格式错误",
	TokenInValid:         "token已过期",

	ERROR: "失败",
}

func GetErrMsg(code int) string {
	return msg[code]
}
