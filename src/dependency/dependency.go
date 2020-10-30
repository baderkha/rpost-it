package dependency

import (
	"rpost-it/src/controller"
	"rpost-it/src/service"
	"rpost-it/src/util"

	"gorm.io/gorm"
)

// ServerEnvDependency : Server dependencies required to build controller dependencies
type ServerEnvDependency struct {
	HashStrength uint   `json:"hashStrength"`
	JWTTokenTime int64  `json:"jwtTokenTimeMinutes"`
	JWTIssuer    string `json:"jwtIssuer"`
	JWTSecret    string `json:"jwtSecret"`
}

// Dependency : Dependencies required for the router
type Dependency struct {
	*controller.AccountController
	*controller.PostController
	*controller.CommunityController
}

func makeService(db *gorm.DB, serverEnv *ServerEnvDependency) service.IFacade {
	return service.MakeFacade(
		db,
		&util.JWTHS265{
			Issuer: serverEnv.JWTIssuer,
			Secret: []byte(serverEnv.JWTSecret),
		},
		serverEnv.JWTTokenTime,
		serverEnv.HashStrength,
		&util.PasswordBcrypt{},
	)
}

// MakeDependencies : Makes the Dependencies into a unified accessible facade
func MakeDependencies(db *gorm.DB, serverEnv *ServerEnvDependency) *Dependency {
	svc := makeService(db, serverEnv)
	return &Dependency{
		AccountController: &controller.AccountController{
			BaseController: controller.BaseController{},
			Service:        svc,
		},
		CommunityController: &controller.CommunityController{
			BaseController: controller.BaseController{},
			Service:        svc,
		},
		PostController: &controller.PostController{
			BaseController: controller.BaseController{},
			Service:        svc,
		},
	}
}
