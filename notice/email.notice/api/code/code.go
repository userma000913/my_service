package code

import "email.notice/model"

var e = make(map[int]string)

var (
	ErrOK     = 0
	ErrParams = 499
	ErrSystem = 500
	ErrSend   = 501
)

func init() {
	e = map[int]string{
		ErrOK:     "操作成功",
		ErrParams: "参数错误",
		ErrSystem: "系统内部错误",
		ErrSend:   "邮件发送失败",
	}
}
func GetApiResponse(code int, msg ...string) model.ApiResult {
	defMsg := "操作失败，请稍后重试"
	if errMsg, ok := e[code]; ok && errMsg != "" {
		defMsg = errMsg
	}

	if len(msg) > 0 {
		defMsg = msg[0]
	}

	return model.ApiResult{DmError: code, ErrorMsg: defMsg}
}

func GetResp(data interface{}, code int) (res model.RespBase, resCode int) {
	res = model.RespBase{ApiResult: GetApiResponse(code), Data: data}
	resCode = code
	return
}
