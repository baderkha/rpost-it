package repository

import (
	"gorm.io/gorm"
)

type BaseRepo struct {
	Context *gorm.DB
}

// MakeBaseRepo : Creates a base repository
func MakeBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{
		Context: db,
	}
}

func (b *BaseRepo) GetContext() *gorm.DB {
	return b.Context
}

func (b *BaseRepo) FindById(id string, model interface{}) bool {
	return b.
		GetContext().
		Where("id=?", id).
		Find(model).
		RowsAffected > 0
}

func (b *BaseRepo) FindAll(id string, model interface{}) bool {
	return b.
		GetContext().
		Find(model).
		RowsAffected > 0
}

func (b *BaseRepo) Create(model interface{}) bool {
	return b.
		GetContext().
		Create(model).
		RowsAffected > 0
}

func (b *BaseRepo) Update(model interface{}) bool {
	return b.
		GetContext().
		Save(model).
		RowsAffected > 0
}

func (b *BaseRepo) DeleteById(id string, model interface{}) bool {
	return b.
		GetContext().
		Where("id=?", id).
		Unscoped().
		Delete(model).
		RowsAffected > 0
}
