package models

import (
	"time"

	"gorm.io/gorm"
)

type NotaTransaksi struct {
	gorm.Model
	PegawaiId        uint
	CustomerId       uint
	KodeTransaksi    string
	NamaTransaksi    string
	TanggalTransaksi time.Time
	TotalBarang      int
	Keterangan       string
}

type NotaTransaksiModel struct {
	db *gorm.DB
}

func NewNotaTransaksiModel(connection *gorm.DB) *NotaTransaksiModel {
	return &NotaTransaksiModel{
		db: connection,
	}
}

func (ntm *NotaTransaksiModel) CreateNotaTransaksi(newData NotaTransaksi) (NotaTransaksi, error) {
	err := ntm.db.Create(&newData).Error
	if err != nil {
		return NotaTransaksi{}, err
	}
	return newData, nil
}

func (ntm *NotaTransaksiModel) DeleteNotaTransaksi(id uint) (NotaTransaksi, error) {
	var nota NotaTransaksi
	err := ntm.db.Where("ID = ?", id).First(&nota).Error
	if err != nil {
		return NotaTransaksi{}, err
	}
	err = ntm.db.Delete(&nota).Error
	if err != nil {
		return NotaTransaksi{}, err
	}

	return nota, nil
}
