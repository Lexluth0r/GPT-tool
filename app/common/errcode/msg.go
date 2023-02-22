package errcode

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "服务异常,请稍后重试",
	InvalidParams: "请求参数错误",
	NotFound:      "找不到资源",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
