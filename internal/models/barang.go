package models

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	ID         uint `gorm:"primaryKey"`
	PegawaiId  uint
	KodeBarang string
	NamaBarang string
	Stok       int
	Harga      int
	Keterangan string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type ResponseBarang struct {
	ID         uint
	KodeBarang string
	NamaBarang string
	Stok       int
	Harga      int
	Keterangan string
}

type BarangModel struct {
	DB *gorm.DB
}

func NewBarangModel(db *gorm.DB) *BarangModel {
	return &BarangModel{
		DB: db,
	}
}

func (bm *BarangModel) AddBarang(barang Barang) (*Barang, error) {
	if err := bm.DB.Create(&barang).Error; err != nil {
		return nil, err
	}
	return &barang, nil
}

func (bm *BarangModel) UpdateBarang(id uint, updatedData Barang) (*Barang, error) {
	barang := &Barang{}
	if err := bm.DB.First(barang, id).Error; err != nil {
		return nil, err
	}

	barang.KodeBarang = updatedData.KodeBarang
	barang.NamaBarang = updatedData.NamaBarang
	barang.Stok = updatedData.Stok
	barang.Harga = updatedData.Harga
	barang.Keterangan = updatedData.Keterangan
	barang.UpdatedAt = time.Now()

	if err := bm.DB.Save(barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (bm *BarangModel) DeleteBarang(id uint) (*Barang, error) {
	barang := &Barang{}
	if err := bm.DB.First(barang, id).Error; err != nil {
		return nil, err
	}

	if err := bm.DB.Delete(barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (bm *BarangModel) FindBarang(pegawaiID uint) (*[]Barang, error) {
	barangs := &[]Barang{}
	if err := bm.DB.Where("pegawai_id = ?", pegawaiID).Find(barangs).Error; err != nil {
		return nil, err
	}
	return barangs, nil
}

func (bm *BarangModel) GetBarang(barangID uint) (*Barang, error) {
	barang := &Barang{}
	if err := bm.DB.Where("ID = ?", barangID).First(barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}
