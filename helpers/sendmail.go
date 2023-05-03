package helpers

import (
	"fmt"
	"net/smtp"
)

func Sendmail(email []string) {
	auth:=smtp.PlainAuth(
		"",
		"ade_epick@outlook.com",
		"Google767",
		"smtp.hotmail.com",
	)
	msg:= "Subject: Reset Password\nThis will reset your password" 
	err:= smtp.SendMail(
		"smtp.hotmail.com:587",
		auth,
		"ade_epick@outlook.com",
		[]string(email),
		[]byte(msg),

	)
	if err != nil {
		fmt.Println(err)
	}
}