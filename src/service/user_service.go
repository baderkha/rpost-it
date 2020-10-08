package service

import (
	"comment-me/src/repository"
	"errors"
	"time"
)

type UserRegistrationDetails struct {
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	DateOfBirth time.Time `json:"dob" binding:"required"`
	AccountID   string
}

type IUserService interface {
	RegisterUser(r *UserRegistrationDetails) (*repository.User, error)
}

type UserService struct {
	Repo repository.IUserRepo
}

func (u *UserService) RegisterUser(r *UserRegistrationDetails) (*repository.User, error) {
	user := repository.User{}
	if r.AccountID == "" {
		return nil, errors.New("400, Expecting the account id to not be a nonvalue")
	}
	_, accountIdExists := u.Repo.FindUserByAccountId(r.AccountID)
	if accountIdExists {
		return nil, errors.New("400, The associated account already exists , this is not a brand new registration")
	}
	user.FirstName = r.FirstName
	user.LastName = r.LastName
	user.DateOfBirth = r.DateOfBirth
	user.AccountId = r.AccountID
	isCreated := u.Repo.CreateUser(&user)
	if isCreated {
		return &user, nil
	}
	return nil, errors.New("500, Could Not Create a User Contact Admin")
}
