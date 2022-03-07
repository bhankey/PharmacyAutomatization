package container

import (
	"time"

	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/emailrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/onetimecodesrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/tokenrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/userrepo"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/middleware"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/userhandler"
	"github.com/bhankey/pharmacy-automatization/internal/service/authservice"
)

func (c *Container) GetUserHandler() *userhandler.UserHandler {
	const key = "UserHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*userhandler.UserHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := userhandler.NewUserHandler(c.getBaseHandler(), c.getUserSrv())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getBaseHandler() *http.BaseHandler {
	const key = "BaseHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*http.BaseHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := http.NewHandler(c.logger)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getUserSrv() *authservice.AuthService {
	const key = "UserSrv"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*authservice.AuthService)
		if ok {
			return typedDependency
		}
	}

	typedDependency := authservice.NewUserService(
		c.getUserStorage(),
		c.getTokenStorage(),
		c.getEmailStorage(),
		c.getOneTimeCodesPasswordStorage(),
		c.passwordSalt,
		c.jwtKey,
	)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getUserStorage() *userrepo.UserRepo {
	const key = "UserStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*userrepo.UserRepo)
		if ok {
			return typedDependency
		}
	}

	typedDependency := userrepo.NewUserRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getTokenStorage() *tokenrepo.TokenRepo {
	const key = "TokenStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*tokenrepo.TokenRepo)
		if ok {
			return typedDependency
		}
	}

	typedDependency := tokenrepo.NewTokenRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getEmailStorage() *emailrepo.EmailRepo {
	const key = "EmailStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*emailrepo.EmailRepo)
		if ok {
			return typedDependency
		}
	}

	typedDependency := emailrepo.NewEmailRepo(c.smtpClient, c.smtpMessageFrom)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getOneTimeCodesPasswordStorage() *onetimecodesrepo.ResetCodesRepo {
	const key = "OneTimeCodesPasswordStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*onetimecodesrepo.ResetCodesRepo)
		if ok {
			return typedDependency
		}
	}

	const timeOfLife = time.Second * 15 // TODO move to config or something else

	typedDependency := onetimecodesrepo.NewResetCodesRepo(c.redisConnection, timeOfLife)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) GetAuthMiddleware() *middleware.AuthMiddleware {
	const key = "AuthMiddleware"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*middleware.AuthMiddleware)
		if ok {
			return typedDependency
		}
	}

	typedDependency := middleware.NewAuthMiddleware(c.logger, c.jwtKey)

	c.dependencies[key] = typedDependency

	return typedDependency
}
