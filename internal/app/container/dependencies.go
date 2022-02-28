package container

import (
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/middleware"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/userhandler"
	"github.com/bhankey/pharmacy-automatization/internal/repository/userrepo"
	"github.com/bhankey/pharmacy-automatization/internal/service/userservice"
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

func (c *Container) getUserSrv() *userservice.UserService {
	const key = "UserSrv"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*userservice.UserService)
		if ok {
			return typedDependency
		}
	}

	typedDependency := userservice.NewUserService(c.getUserStorage(), c.passwordSalt, c.jwtKey)

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
