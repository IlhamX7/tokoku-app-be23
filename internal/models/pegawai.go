package models

import (
	"time"

	"gorm.io/gorm"
)

type Pegawai struct {
	gorm.Model
	Username      string
	Password      string
	Email         string
	AdminId       uint
	Customer      []Customer      `gorm:"foreignKey:PegawaiId"`
	Barang        []Barang        `gorm:"foreignKey:PegawaiId"`
	NotaTransaksi []NotaTransaksi `gorm:"foreignKey:PegawaiId"`
}

type PegawaiModel struct {
	db *gorm.DB
}

func NewPegawaiModel(connection *gorm.DB) *PegawaiModel {
	return &PegawaiModel{
		db: connection,
	}
}

func (pm *PegawaiModel) AddPegawai(newData Pegawai) (Pegawai, error) {
	err := pm.db.Create(&newData).Error
	if err != nil {
		return Pegawai{}, err
	}
	return newData, nil
}

func (pm *PegawaiModel) UpdatePegawai(id uint, newUsername string, newPassword string, newEmail string) (Pegawai, error) {
	var pegawai Pegawai
	err := pm.db.Where("ID = ?", id).First(&pegawai).Error
	if err != nil {
		return Pegawai{}, err
	}
	pegawai.Username = newUsername
	pegawai.Password = newPassword
	pegawai.Email = newEmail
	pegawai.UpdatedAt = time.Now()

	err = pm.db.Save(&pegawai).Error
	if err != nil {
		return Pegawai{}, err
	}

	return pegawai, nil
}

func (pm *PegawaiModel) DeletePegawai(id uint) (Pegawai, error) {
	var pegawai Pegawai
	err := pm.db.Where("ID = ?", id).First(&pegawai).Error
	if err != nil {
		return Pegawai{}, err
	}
	err = pm.db.Delete(&pegawai).Error
	if err != nil {
		return Pegawai{}, err
	}

	return pegawai, nil
}

func (pm *PegawaiModel) FindPegawai(id uint) ([]Pegawai, error) {
	var pegawai []Pegawai
	err := pm.db.Where("Admin_Id = ?", id).Find(&pegawai).Error
	if err != nil {
		return nil, err
	}

	return pegawai, nil
}

func (pm *PegawaiModel) LoginPegawai(username string, password string) (Pegawai, error) {
	var result Pegawai
	err := pm.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return Pegawai{}, err
	}
	return result, nil
}
