package repository

import (
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Title        string // the main Title for the viewer
	Description  string // Just the bio , can be long text
	UniqueId     string `gorm:"type:varchar(100)"` // CodWarzone forexample , must be unique
	AccountOwner uint 
}

type ICommunityRepo interface {
}
