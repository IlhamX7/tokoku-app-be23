package models

import (
	"gorm.io/gorm"
)

type DetailTransaksi struct {
	gorm.Model
	NotaTransaksiID uint
	BarangID        uint
	JumlahBarang    int
	HargaBarang     int
}

type DetailTransaksiModel struct {
	db *gorm.DB
}

func NewDetailTransaksiModel(connection *gorm.DB) *DetailTransaksiModel {
	return &DetailTransaksiModel{
		db: connection,
	}
}

func (dtm *DetailTransaksiModel) AddDetailTransaksi(detail DetailTransaksi) (*DetailTransaksi, error) {
	if err := dtm.db.Create(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

func (dtm *DetailTransaksiModel) DeleteDetailTransaksi(id uint) (*DetailTransaksi, error) {
	detail := &DetailTransaksi{}
	if err := dtm.db.First(detail, id).Error; err != nil {
		return nil, err
	}

	if err := dtm.db.Delete(detail).Error; err != nil {
		return nil, err
	}
	return detail, nil
}

func (dtm *DetailTransaksiModel) FindDetailTransaksi(notaTransaksiID uint) ([]DetailTransaksi, error) {
	var details []DetailTransaksi
	err := dtm.db.Where("nota_transaksi_id = ?", notaTransaksiID).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}
