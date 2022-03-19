package container

import (
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/addressrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/pharmacyrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/productrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/receiptrepo"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/pharmacyhandler"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/purchasehandler"
	"github.com/bhankey/pharmacy-automatization/internal/service/pharmacyservice"
	"github.com/bhankey/pharmacy-automatization/internal/service/purchaseservice"
	"time"

	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/emailrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/onetimecodesrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/tokenrepo"
	"github.com/bhankey/pharmacy-automatization/internal/adapter/repository/userrepo"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/middleware"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/authhandler"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/swaggerhandler"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/userhandler"
	"github.com/bhankey/pharmacy-automatization/internal/service/authservice"
	"github.com/bhankey/pharmacy-automatization/internal/service/userservice"
)

func (c *Container) GetV1AuthHandler() *authhandler.AuthHandler {
	const key = "V1AuthHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*authhandler.AuthHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := authhandler.NewAuthHandler(c.getBaseHandler(), c.getAuthSrv())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) GetV1SwaggerHandler() *swaggerhandler.SwaggerHandler {
	const key = "V1SwaggerHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*swaggerhandler.SwaggerHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := swaggerhandler.NewSwaggerHandler(c.getBaseHandler())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) GetV1UserHandler() *userhandler.UserHandler {
	const key = "V1UserHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*userhandler.UserHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := userhandler.NewUserHandler(c.getBaseHandler(), c.getUserSrv(), c.GetAuthMiddleware())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) GetV1PharmacyHandler() *pharmacyhandler.Handler {
	const key = "V1PharmacyHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*pharmacyhandler.Handler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := pharmacyhandler.NewPharmacyHandler(c.getBaseHandler(), c.getPharmacyService(), c.GetAuthMiddleware())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) GetV1PurchaseHandler() *purchasehandler.Handler {
	const key = "V1PurchaseHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*purchasehandler.Handler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := purchasehandler.NewPurchaseHandler(c.getBaseHandler(), c.getPurchaseService(), c.GetAuthMiddleware())

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

	typedDependency := userservice.NewUserService(
		c.getUserStorage(),
		c.getEmailStorage(),
		c.getOneTimeCodesPasswordStorage(),
	)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getAuthSrv() *authservice.AuthService {
	const key = "AuthSrv"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*authservice.AuthService)
		if ok {
			return typedDependency
		}
	}

	typedDependency := authservice.NewAuthService(
		c.getUserStorage(),
		c.getTokenStorage(),
		c.jwtKey,
	)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getUserStorage() *userrepo.Repository {
	const key = "UserStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*userrepo.Repository)
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

func (c *Container) getAddressStorage() *addressrepo.Repository {
	const key = "AddressStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*addressrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := addressrepo.NewAddressRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getProductStorage() *productrepo.Repository {
	const key = "ProductStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*productrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := productrepo.NewProductRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getReceiptStorage() *receiptrepo.Repository {
	const key = "ProductStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*receiptrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := receiptrepo.NewReceiptRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getPharmacyStorage() *pharmacyrepo.Repository {
	const key = "PharmacyStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*pharmacyrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := pharmacyrepo.NewPharmacyRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getPharmacyService() *pharmacyservice.Service {
	const key = "PharmacyService"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*pharmacyservice.Service)
		if ok {
			return typedDependency
		}
	}

	typedDependency := pharmacyservice.NewPharmacyService(c.getPharmacyStorage(), c.getAddressStorage())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getPurchaseService() *purchaseservice.Service {
	const key = "PurchaseService"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*purchaseservice.Service)
		if ok {
			return typedDependency
		}
	}

	typedDependency := purchaseservice.NewPurchaseService(c.getProductStorage(), c.getReceiptStorage())

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
