package repository

import (
	"gorm.io/gorm"
)

// Community : Community model
type Community struct {
	gorm.Model
	Title        string // the main Title for the viewer
	Description  string // Just the bio , can be long text
	UniqueID     string `gorm:"type:varchar(100)"` // CodWarzone forexample , must be unique
	AccountOwner uint
}

// ICommunityRepo : Community repository
type ICommunityRepo interface {
	// FindCommunityByID : finds a specific community by the primary key
	FindCommunityByID(id string) (*Community, bool)
	// FindCommunitiesByAccountOwner : Finds a community by account owner / account id
	FindCommunitiesByAccountOwner(accountID string) *[]Community
	// FindCommunityByUniqueID : Finds a community by the unique id for that community ie codWarzone
	FindCommunityByUniqueID(uniqueID string) (*Community, bool)
	// FindCommunityByUniqueID : Finds a community by the unique id for that community ie codWarzone
	CreateCommunity(com *Community) bool
	// DeleteCommunityByIDAndAccountOwner : Deletes a Community by a specific id and the account owner , to ensure saftey of delete operation
	DeleteCommunityByIDAndAccountOwner(id string, AccountOwner string) bool
}

// CommunityRepo : Encapsulates the data store for the community
type CommunityRepo struct {
	BaseRepo
}

// FindCommunityByID : finds a specific community by the primary key
func (c *CommunityRepo) FindCommunityByID(id string) (*Community, bool) {
	var model Community
	isFound := c.FindById(id, &model)
	return &model, isFound
}

// FindCommunitiesByAccountOwner : Finds a community by account owner / account id
func (c *CommunityRepo) FindCommunitiesByAccountOwner(accountID string) *[]Community {
	var model []Community
	c.
		GetContext().
		Where("account_owner=?", accountID).
		Find(&model)
	return &model
}

// FindCommunityByUniqueID : Finds a community by the unique id for that community ie codWarzone
func (c *CommunityRepo) FindCommunityByUniqueID(uniqueID string) (*Community, bool) {
	var model Community
	isFound := c.
		GetContext().
		Where("unique_id=?", uniqueID).
		First(&model).
		RowsAffected > 0
	return &model, isFound
}

// CreateCommunity : Creates a community record in the db
func (c *CommunityRepo) CreateCommunity(com *Community) bool {
	return c.Create(com)
}

// DeleteCommunityByIDAndAccountOwner : Deletes a Community by a specific id and the account owner , to ensure saftey of delete operation
func (c *CommunityRepo) DeleteCommunityByIDAndAccountOwner(id string, accountOwner string) bool {
	isDeleted := c.
		GetContext().
		Where("unique_id=?", id).
		Where("account_owner=?", accountOwner).
		Unscoped().
		Delete(&Community{}).
		RowsAffected > 0
	return isDeleted
}
