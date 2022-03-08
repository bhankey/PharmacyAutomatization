package emailrepo

import (
	"bytes"
	"fmt"
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
	errBase := fmt.Sprintf("emailrepo.SendResetPasswordCode(%s, %s)", email, code)

	htmlTemplate, err := template.ParseFiles(
		"./internal/adapter/repository/emailrepo/reset_password.html",
	) // TODO really bad
	if err != nil {
		return fmt.Errorf("%s: failed to parse html template: %w", errBase, err)
	}

	var body bytes.Buffer
	if err := htmlTemplate.Execute(&body, struct {
		Code string
	}{
		Code: code,
	}); err != nil {
		return fmt.Errorf("%s: failed to execute html template: %w", errBase, err)
	}

	emailMessage := mail.NewMSG()

	emailMessage.SetFrom(r.from).
		AddTo(email).
		SetSubject("Reset password").
		SetBody(mail.TextHTML, body.String())

	if err := emailMessage.Send(r.smtp); err != nil {
		return fmt.Errorf("%s: failed to send message: %w", errBase, err)
	}

	if emailMessage.Error != nil {
		return fmt.Errorf("%s: failed to send message email error: %w", errBase, emailMessage.Error)
	}

	return nil
}
