package container

import (
	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	masterPostgresDB *sqlx.DB
	slavePostgresDB  *sqlx.DB
	redisConnection  *redis.Client
	logger           logger.Logger

	dependencies map[string]interface{}
}

func NewContainer(log logger.Logger, masterPostgres, slavePostgres *sqlx.DB, redis *redis.Client) *Container {
	return &Container{
		masterPostgresDB: masterPostgres,
		slavePostgresDB:  slavePostgres,
		redisConnection:  redis,
		logger:           log,
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
