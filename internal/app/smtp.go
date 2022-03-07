package app

import (
	"time"

	"github.com/bhankey/pharmacy-automatization/internal/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

func newSMTPClient(c config.Config) (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()

	server.Host = c.SMTP.Host
	server.Port = c.SMTP.Port
	server.Username = c.SMTP.User
	server.Password = c.SMTP.Password
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = true
	server.ConnectTimeout = 10 * time.Second // nolint: gomnd
	server.SendTimeout = 10 * time.Second    // nolint: gomnd
	// TODO delete Debug: server.TLSConfig = &tls.Config{InsecureSkipVerify: true} // nolint:

	client, err := server.Connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}
