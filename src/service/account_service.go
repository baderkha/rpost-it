package service

import (
	"errors"
	"fmt"
	"rpost-it/src/repository"
	"rpost-it/src/util"
	"strings"
	"time"
)

// RegistrationDetails : things we expected for registration input
type RegistrationDetails struct {
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	DateOfBirth time.Time `json:"dob" binding:"required"`
	AvatarID    string    `json:"avatarId" binding:"required"`
	Password    string    `json:"password" binding:"required"`
}

// LoginDetails  : Things we expect when loggin in
type LoginDetails struct {
	AvatarID string `json:"avatarId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AccessRequest :
type AccessRequest struct {
	AccountIDAccess string
	Operation       string
}

// RoleAccess : still thinking about this one
type RoleAccess struct {
	Resource string
	Verb     string
}

// JWT  : base jwt repoonse type
type JWT struct {
	Token       string              `json:"token"`
	ExpiryEpoch int64               `json:"expiryEpoch"`
	Account     *repository.Account `json:"account"`
}

// JWTClaim : what we want the claim to have
type JWTClaim struct {
	AvatarID  string `json:"avatarId"`
	AccountID string `json:"accountId"`
}

// AccountService : Account service for major auth logic
type AccountService struct {
	Repo                   repository.IAccountRepo
	JWTHelper              util.IJwtHelper
	PasswordHelper         util.IPassword
	PassWordHashedStrength uint
	JWTValidityMinutes     int64
}

// obfuscateAccount : This is a very important method to hide the password ,please use it if you intend on returning
func (a *AccountService) obfuscateAccountTrustedUser(account *repository.Account) *repository.Account {
	account.Password = ""
	return account
}

// RegisterAccount : Register an account to the database
func (a *AccountService) RegisterAccount(r *RegistrationDetails) (*repository.Account, error) {
	_, accountAlreadyExists := a.Repo.FindByAvatarIdOrByEmail(r.AvatarID, r.Email)
	if accountAlreadyExists {
		return nil, errors.New("400, This account already exists")
	}
	acc := repository.Account{}
	acc.AvatarId = r.AvatarID
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

// LoginAccount : Login into the account and return back a jwt
func (a *AccountService) LoginAccount(l *LoginDetails) (*JWT, error) {
	acc, exists := a.Repo.FindByAvatarId(l.AvatarID)
	if !exists {
		return nil, errors.New("401, Invalid Avatar ID or password")
	}
	isValidPassword := a.PasswordHelper.CheckPasswordHash(l.Password, acc.Password)
	if isValidPassword {
		return a.generateJWTForValidAccount(acc)
	}
	return nil, errors.New("401, Invalid Avatar ID or password")
}

// DeleteServiceByIDInternalOnly : internal delete for account , not to be used outside
func (a *AccountService) DeleteServiceByIDInternalOnly(id string) bool {
	// guard check
	if id == "" {
		return false
	}
	return a.Repo.DeleteAccountById(id)
}

func (a *AccountService) generateJWTForValidAccount(acc *repository.Account) (*JWT, error) {
	token, validity, err := a.JWTHelper.GenerateWebToken(fmt.Sprintf("%d", acc.ID), a.JWTValidityMinutes)
	if err != nil {
		return nil, errors.New("500, Could Not Generate JWT contact api admin")
	}
	return &JWT{
		Token:       token,
		ExpiryEpoch: validity,
		Account:     a.obfuscateAccountTrustedUser(acc),
	}, nil
}

// RefreshJWTToken : refresh a jwt token for an account Id
func (a *AccountService) RefreshJWTToken(accountId string) (*JWT, error) {
	if accountId == "" {
		return nil, errors.New("Expecting an account ID")
	}
	acc, isFound := a.Repo.FindByAccountId(accountId)
	if !isFound {
		return nil, errors.New("400, Account id not found")
	}
	return a.generateJWTForValidAccount(acc)
}

// ValidateJWT : ensure the jwt is valid
// this needs a caching repository ?
func (a *AccountService) ValidateJWT(roleAccess *RoleAccess, jwt string) (*repository.Account, error) {
	isValid, claims := a.JWTHelper.ValidateWebToken(jwt)
	if !isValid {
		return nil, errors.New("401, Not a Valid JWT")
	}
	acc, _ := a.Repo.FindByAccountId(claims.Subject)
	return a.obfuscateAccountTrustedUser(acc), nil
}

// ValidateAccountExists : check if account exists for an accountID
func (a *AccountService) ValidateAccountExists(accountID string) bool {
	_, isFound := a.Repo.FindByAccountId(accountID)
	return isFound
}

// GetAccountInfoByJWT : fetch a valid account by the jwt token , this will also check the validity / expirtation of token
func (a *AccountService) GetAccountInfoByJWT(JWTBearer string) (*repository.Account, error) {
	if JWTBearer == "" {
		return nil, errors.New("400, Missing the JWT ")
	}
	splitBearer := strings.Split(JWTBearer, "Bearer ")
	if len(splitBearer) <= 1 {
		return nil, errors.New("Bearer is Supposed to included in the request")
	}
	JWTBearer = splitBearer[1]
	isValid, claims := a.JWTHelper.ValidateWebToken(JWTBearer)
	if !isValid {
		return nil, errors.New("401, Not a Valid JWT")
	}
	acc, isFound := a.Repo.FindByAccountId(claims.Subject)
	if !isFound {
		return nil, errors.New("404 , This account was not found")
	}
	return a.obfuscateAccountTrustedUser(acc), nil
}
