package app

import (
	"fmt"
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
	// Debug: server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	client, err := server.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect smtp server: %w", err)
	}

	return client, nil
}
