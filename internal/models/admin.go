package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string
	Password string
}

type AdminModel struct {
	db *gorm.DB
}

func NewAdminModel(connection *gorm.DB) *AdminModel {
	return &AdminModel{
		db: connection,
	}
}
