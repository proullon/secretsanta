package bucket

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"
)

const GmailServer = "smtp.gmail.com"
const GmailPort = "587"

type smtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

type EmailAccount struct {
	Email    string
	Password string
	Server   string
	Port     string

	From    string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
Content-Type: text/html;


<html>
<head>
</head>
<body>

{{.Body}}

</body>
</html>
`

func templateMessage(from string, to string, subject string, body string) (doc bytes.Buffer, err error) {

	context := &smtpTemplateData{from, to, subject, body}
	t := template.New("emailTemplate")
	t, err = t.Parse(emailTemplate)
	if err != nil {
		log.Printf("error trying to parse mail template: %s\n", err)
		return
	}
	err = t.Execute(&doc, context)
	if err != nil {
		log.Printf("error trying to execute mail template: %s\n", err)
	}

	return
}

func sendEmail(conf EmailAccount, subject string, body string, receiver string) (err error) {

	doc, err := templateMessage(conf.From, receiver, subject, body)
	if err != nil {
		log.Printf("SendEmail: Cannot template message: %s\n", err)
		return
	}

	auth := smtp.PlainAuth("", conf.Email, conf.Password, conf.Server)
	err = smtp.SendMail(conf.Server+":"+conf.Port, auth, conf.Email, []string{receiver}, doc.Bytes())
	if err != nil {
		log.Printf("SendEmail: Cannot send mail : %s\n", err)
	}
	return
}
