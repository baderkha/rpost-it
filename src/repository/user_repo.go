package repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(255);"`
	LastName    string `gorm:"type:varchar(255);"`
	DateOfBirth time.Time
	AccountId   string `gorm:"account_id"`
}

type IUserRepo interface {
	FindUserByAccountId(accountId string) (*User, bool)
	FindByUserId(id string) (*User, bool)
	CreateUser(user *User) bool
	DeleteUserById(id string) bool
}

type UserRepo struct {
	BaseRepo
}

func (u *UserRepo) FindByUserId(id string) (*User, bool) {
	var user User
	db := u.GetContext()
	isFound := db.Where("id=?", id).First(&user).RowsAffected > 0
	return &user, isFound
}

func (u *UserRepo) FindUserByAccountId(id string) (*User, bool) {
	db := u.GetContext()
	var user User
	isFound := db.Where("account_id=?", id).First(&user).RowsAffected > 0
	return &user, isFound
}

func (u *UserRepo) CreateUser(user *User) bool {
	return u.BaseRepo.Create(user)
}

func (u *UserRepo) DeleteUserById(id string) bool {
	return u.BaseRepo.DeleteById(id, &User{})
}
