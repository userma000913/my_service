package api

import (
	"email.notice/api/code"
	"email.notice/email"
	"email.notice/model"
	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	var req model.EmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// todo
		return
	}

	// 参数校验

	// 发送邮件
	rCode := email.SendEmail(req.To, req.Subject, req.Body)

	// 响应消息
	code.GetResp(nil, rCode)
	return
}
