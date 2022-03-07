package container

import (
	"github.com/bhankey/pharmacy-automatization/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Container struct {
	masterPostgresDB *sqlx.DB
	slavePostgresDB  *sqlx.DB
	redisConnection  *redis.Client
	smtpClient       *mail.SMTPClient
	logger           logger.Logger

	passwordSalt    string
	jwtKey          string
	smtpMessageFrom string

	dependencies map[string]interface{}
}

func NewContainer(
	log logger.Logger,
	masterPostgres, slavePostgres *sqlx.DB,
	redis *redis.Client,
	smtpClient *mail.SMTPClient,
	passwordSalt,
	jwtKey,
	smtpMessageFrom string,
) *Container {
	return &Container{
		masterPostgresDB: masterPostgres,
		slavePostgresDB:  slavePostgres,
		redisConnection:  redis,
		smtpClient:       smtpClient,
		logger:           log,
		passwordSalt:     passwordSalt,
		jwtKey:           jwtKey,
		smtpMessageFrom:  smtpMessageFrom,
		dependencies:     make(map[string]interface{}),
	}
}

func (c *Container) CloseAllConnections() {
	if err := c.masterPostgresDB.Close(); err != nil {
		c.logger.Errorf("failed to close master postgres connection error: %v", err)
	}

	if err := c.slavePostgresDB.Close(); err != nil {
		c.logger.Errorf("failed to close slave postgres connection error: %v", err)
	}

	if err := c.redisConnection.Close(); err != nil {
		c.logger.Errorf("failed to close redis connection error: %v", err)
	}
}
