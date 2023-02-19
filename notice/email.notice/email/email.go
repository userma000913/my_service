package email

import (
	"crypto/tls"
	"email.notice/api/code"
	"email.notice/config"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

func SendEmail(To, subject string, body string) int {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

func send(to []string, subject string, body string) int {
	from := config.EmailConf.Email.From
	nickname := config.EmailConf.Email.Nickname
	secret := config.EmailConf.Email.Secret
	host := config.EmailConf.Email.Host
	port := config.EmailConf.Email.Port
	isSSL := config.EmailConf.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
		if err != nil {
			return code.ErrSend
		}
	} else {
		err = e.Send(hostAddr, auth)
		if err != nil {
			return code.ErrSend
		}
	}
	return code.ErrOK
}
