package dependency

import (
	"comment-me/src/controller"
	"comment-me/src/repository"
	"comment-me/src/service"
	"comment-me/src/util"

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

// MakeDependencies : Makes the Dependencies into a unified accessible facade
func MakeDependencies(db *gorm.DB, serverEnv *ServerEnvDependency) *Dependency {
	return &Dependency{
		AccountController: &controller.AccountController{
			BaseController: controller.BaseController{},
			Service: &service.AccountService{
				JWTValidityMinutes: serverEnv.JWTTokenTime,
				JWTHelper: &util.JWTHS265{
					Issuer: serverEnv.JWTIssuer,
					Secret: []byte(serverEnv.JWTSecret),
				},
				PassWordHashedStrength: serverEnv.HashStrength,
				PasswordHelper:         &util.PasswordBcrypt{},
				Repo: &repository.AccountRepo{
					BaseRepo: repository.BaseRepo{
						Context: db,
					},
				},
				UserService: &service.UserService{
					Repo: &repository.UserRepo{
						BaseRepo: repository.BaseRepo{
							Context: db,
						},
					},
				},
			},
		},
		CommunityController: &controller.CommunityController{
			BaseController: controller.BaseController{},
			Service: &service.CommunityService{
				AccountService: &service.AccountService{
					JWTValidityMinutes: serverEnv.JWTTokenTime,
					JWTHelper: &util.JWTHS265{
						Issuer: serverEnv.JWTIssuer,
						Secret: []byte(serverEnv.JWTSecret),
					},
					PassWordHashedStrength: serverEnv.HashStrength,
					PasswordHelper:         &util.PasswordBcrypt{},
					Repo: &repository.AccountRepo{
						BaseRepo: repository.BaseRepo{
							Context: db,
						},
					},
					UserService: &service.UserService{
						Repo: &repository.UserRepo{
							BaseRepo: repository.BaseRepo{
								Context: db,
							},
						},
					},
				},
				Repo: &repository.CommunityRepo{
					BaseRepo: repository.BaseRepo{
						Context: db,
					},
				},
			},
		},
		PostController: &controller.PostController{
			BaseController: controller.BaseController{},
			Service: &service.PostService{
				Repo: &repository.PostRepo{
					BaseRepo: repository.BaseRepo{
						Context: db,
					},
				},
				CommunityService: &service.CommunityService{
					Repo: &repository.CommunityRepo{
						BaseRepo: repository.BaseRepo{
							Context: db,
						},
					},
				},
				AccountService: &service.AccountService{
					JWTValidityMinutes: serverEnv.JWTTokenTime,
					JWTHelper: &util.JWTHS265{
						Issuer: serverEnv.JWTIssuer,
						Secret: []byte(serverEnv.JWTSecret),
					},
					PassWordHashedStrength: serverEnv.HashStrength,
					PasswordHelper:         &util.PasswordBcrypt{},
					Repo: &repository.AccountRepo{
						BaseRepo: repository.BaseRepo{
							Context: db,
						},
					},
					UserService: &service.UserService{
						Repo: &repository.UserRepo{
							BaseRepo: repository.BaseRepo{
								Context: db,
							},
						},
					},
				},
			},
		},
	}
}
