package router

import (
	"email.notice/api"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	e := r.Group("/email")
	e.POST("/send_emil", api.SendEmail)
}
