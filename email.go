package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendReminder(t *template.Template, data Data) error {

	var (
		smtpAddr = fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
		auth     = smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
		tmp      bytes.Buffer
	)

	if err := t.Execute(&tmp, data); err != nil {
		return err
	}

	var message = buildMessage(data.Email, os.Getenv("SUBJECT"), tmp.Bytes())

	if err := smtp.SendMail(smtpAddr, auth, os.Getenv("SMTP_EMAIL"), []string{data.Email}, message); err != nil {
		fmt.Println("disini, ", err.Error())
		return err
	}

	return nil
}

func buildMessage(to, subject string, template []byte) []byte {

	var msg bytes.Buffer
	crlf := "\r\n"
	msg.WriteString(fmt.Sprintf("From:%s <%s>", os.Getenv("SMTP_EMAIL_NAME"), os.Getenv("SMTP_EMAIL")))
	msg.WriteString(crlf)
	msg.WriteString("To:" + to)
	msg.WriteString(crlf)
	msg.WriteString(fmt.Sprintf("Subject:%s", subject))
	msg.WriteString(crlf)
	msg.WriteString("MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n")
	msg.Write(template)
	msg.WriteString(crlf)

	return msg.Bytes()
}
