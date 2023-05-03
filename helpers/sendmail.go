package helpers

import (
	"fmt"
	"log"
	"net/smtp"
)

func Sendmail(email []string) {
	
	auth:=smtp.PlainAuth(
		"",
		"ade_epick@outlook.com",
		"Google767",
		"smtp.live.com",
	)
	msg:= "Subject: Reset Password\nThis will reset your password" 
	err:= smtp.SendMail(
		"smtp.live.com:587",
		auth,
		"ade_epick@outlook.com",
		[]string(email),
		[]byte(msg),

	)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		
		
	}
}