package repository

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	User     User   `gorm:"foreignKey:account_id"`
	AvatarId string `gorm:"type:varchar(255);"`
	Password string `gorm:"type:varchar(255);"`
	Email    string `gorm:"type:varchar(255);"`
	IsValid  bool
}

type IAccountRepo interface {
	FindByAvatarId(avatarId string) (*Account, bool)
	FindByAccountId(id string) (*Account, bool)
	FindByAvatarIdOrByEmail(avatarId string, email string) (*Account, bool)
	CreateAccount(acc *Account) bool
	DeleteAccountById(id string) bool
}

type AccountRepo struct {
	BaseRepo
}

func (a *AccountRepo) getJoinedContext() *gorm.DB {
	return a.
		GetContext().
		Joins("User")

}

func (a *AccountRepo) FindByAvatarId(avatarId string) (*Account, bool) {
	var acc Account
	db := a.getJoinedContext()
	isFound := db.Where("avatar_id=?", avatarId).First(&acc).RowsAffected > 0
	return &acc, isFound
}

func (a *AccountRepo) FindByAvatarIdOrByEmail(avatarId string, email string) (*Account, bool) {
	var acc Account
	db := a.getJoinedContext()
	isFound := db.
		Where("avatar_id=?", avatarId).
		Or("email=?", email).
		First(&acc).
		RowsAffected > 0
	return &acc, isFound
}

func (a *AccountRepo) CreateAccount(acc *Account) bool {
	return a.Create(acc)
}

func (a *AccountRepo) FindByAccountId(id string) (*Account, bool) {
	var acc Account
	db := a.getJoinedContext()
	isFound := db.
		Where("accounts.id=?", id).
		First(&acc).
		RowsAffected > 0
	return &acc, isFound
}

func (a *AccountRepo) DeleteAccountById(id string) bool {
	return a.DeleteById(id, &Account{})
}
