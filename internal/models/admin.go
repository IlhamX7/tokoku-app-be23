package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	ID       uint `gorm:"primarykey"`
	Username string
	Password string
	Pegawai  []Pegawai `gorm:"foreignKey:AdminId"`
}

type AdminModel struct {
	db *gorm.DB
}

func NewAdminModel(connection *gorm.DB) *AdminModel {
	return &AdminModel{
		db: connection,
	}
}

func (am *AdminModel) CreateAdmin(admin Admin) (Admin, error) {
	err := am.db.Save(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}

func (am *AdminModel) LoginAdmin(username string, password string) (Admin, error) {
	var result Admin
	err := am.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return Admin{}, err
	}
	return result, nil
}
