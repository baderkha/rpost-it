package service

import (
	"comment-me/src/repository"
	"comment-me/src/util"
	"errors"
	"fmt"
	"time"
)

type RegistrationDetails struct {
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	DateOfBirth time.Time `json:"dob" binding:"required"`
	AvatarId    string    `json:"avatarId" binding:"required"`
	Password    string    `json:"password" binding:"required"`
}

type LoginDetails struct {
	AvatarId string `json:"avatarId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccessRequest struct {
	AccountIdAccess string
	Operation       string
}

type RoleAccess struct {
	Resource string
	Verb     string
}

type JWT struct {
	Token       string `json:"token"`
	ExpiryEpoch int64  `json:"expiryEpoch"`
}

type JWTClaim struct {
	AvatarId  string `json:"avatarId"`
	AccountId string `json:"accountId"`
}

type IAccountService interface {
	RegisterAccount(r *RegistrationDetails) (*repository.Account, error)
	RegisterAccountAndUser(r *RegistrationDetails) (*repository.Account, error)
	LoginAccount(l *LoginDetails) (*JWT, error)
	ValidateJWT(r *RoleAccess, jwt string) error
}

type AccountService struct {
	Repo                   repository.IAccountRepo
	UserService            IUserService
	JWTHelper              util.IJwtHelper
	PasswordHelper         util.IPassword
	PassWordHashedStrength uint
	JWTValidityMinutes     int64
}

func (a *AccountService) RegisterAccount(r *RegistrationDetails) (*repository.Account, error) {
	_, accountAlreadyExists := a.Repo.FindByAvatarIdOrByEmail(r.AvatarId, r.Email)
	if accountAlreadyExists {
		return nil, errors.New("400, This account already exists")
	}
	acc := repository.Account{}
	acc.AvatarId = r.AvatarId
	hashedPass, err := a.PasswordHelper.HashPassword(r.Password, int(a.PassWordHashedStrength))
	if err != nil {
		return nil, errors.New("500,Could Not Create Profile. Please Contact API Admin")
	}
	acc.Password = hashedPass
	acc.Email = r.Email
	isCreated := a.Repo.CreateAccount(&acc)
	if isCreated {
		return &acc, nil
	}
	return nil, errors.New("500,Could Not Create Profile. Please Contact API Admin")
}

func (a *AccountService) RegisterAccountAndUser(r *RegistrationDetails) (*repository.Account, error) {
	acc, err := a.RegisterAccount(r)
	if err != nil {
		return nil, err
	}
	user, err := a.UserService.RegisterUser(&UserRegistrationDetails{
		AccountID:   fmt.Sprintf("%d", acc.ID),
		DateOfBirth: r.DateOfBirth,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
	})
	if err != nil {
		_ = a.Repo.DeleteAccountById(fmt.Sprintf("%d", acc.ID))
		return nil, err
	}
	acc.User = *user
	acc, _ = a.Repo.FindByAccountId(fmt.Sprintf("%d", acc.ID))
	return acc, nil
}

func (a *AccountService) LoginAccount(l *LoginDetails) (*JWT, error) {
	acc, exists := a.Repo.FindByAvatarId(l.AvatarId)
	if !exists {
		return nil, errors.New("401, Invalid Avatar ID or password")
	}
	isValidPassword := a.PasswordHelper.CheckPasswordHash(l.Password, acc.Password)
	if isValidPassword {
		token, validity, err := a.JWTHelper.GenerateWebToken(fmt.Sprintf("%d", acc.ID), a.JWTValidityMinutes)
		if err != nil {
			return nil, errors.New("500, Could Not Generate JWT contact api admin")
		}
		return &JWT{
			Token:       token,
			ExpiryEpoch: validity,
		}, nil
	}
	return nil, errors.New("401, Invalid Avatar ID or password")
}

func (a *AccountService) ValidateJWT(roleAccess *RoleAccess, jwt string) error {
	isValid, _ := a.JWTHelper.ValidateWebToken(jwt)
	if !isValid {
		return errors.New("401, Not a Valid JWT")
	}
	return nil
}
