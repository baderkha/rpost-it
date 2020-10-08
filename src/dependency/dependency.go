package dependency

import (
	"comment-me/src/controller"
	"comment-me/src/repository"
	"comment-me/src/service"
	"comment-me/src/util"
	"gorm.io/gorm"
)

type ServerEnvDependency struct {
	HashStrength uint   `json:"hashStrength"`
	JWTTokenTime int64  `json:"jwtTokenTimeMinutes"`
	JWTIssuer    string `json:"jwtIssuer"`
	JWTSecret    string `json:"jwtSecret"`
}

type Dependency struct {
	*controller.AccountController
	*controller.PostController
}

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
		PostController: &controller.PostController{
			BaseController: controller.BaseController{},
			Service: &service.PostService{
				Repo: &repository.PostRepo{
					BaseRepo: repository.BaseRepo{
						Context: db,
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
