package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendEmail(payload ContactUsPayload) {
	conf, err := LoadConfig(".")

	if err != nil {
		fmt.Println("Error occured while loading config", err)
	}

	recipients := []string{"aakash.tyagi2104@gmail.com", "onlinejudge95@gmail.com"}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: New Lead  \n%s\n\n", mimeHeaders)))

	TemplateContext := map[string]string{
		"Name":          payload.Name,
		"Email":         payload.Email,
		"ContactNumber": payload.ContactNumber,
		"UserType":      payload.UserType,
		"Message":       payload.Message,
	}
	tpl, _ := template.ParseFiles("templates/contact_us.html")
	tpl.Execute(&body, TemplateContext)

	auth := smtp.PlainAuth("", conf.Sender, conf.Password, conf.Host)
	err = smtp.SendMail(conf.Host+":"+conf.Port, auth, conf.Sender, recipients, body.Bytes())
	if err != nil {
		return
	}
	fmt.Println("Email Sent!")
}
