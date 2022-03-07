package emailrepo

import (
	"bytes"
	"text/template"

	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailRepo struct {
	from string

	smtp *mail.SMTPClient
}

func NewEmailRepo(smtp *mail.SMTPClient, from string) *EmailRepo {
	return &EmailRepo{
		from: from,
		smtp: smtp,
	}
}

func (r *EmailRepo) SendResetPasswordCode(email string, code string) error {
	htmlTemplate, err := template.ParseFiles("./internal/adapter/repository/emailrepo/reset_password.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := htmlTemplate.Execute(&body, struct {
		Code string
	}{
		Code: code,
	}); err != nil {
		return err
	}

	emailMessage := mail.NewMSG()

	emailMessage.SetFrom(r.from).
		AddTo(email).
		SetSubject("Reset password").
		SetBody(mail.TextHTML, body.String())

	if err := emailMessage.Send(r.smtp); err != nil {
		return err
	}

	if emailMessage.Error != nil {
		return emailMessage.Error
	}

	return nil
}
