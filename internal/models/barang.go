package models

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	PegawaiId       uint
	KodeBarang      string
	NamaBarang      string
	Stok            int
	Harga           int
	Keterangan      string
	DetailTransaksi []DetailTransaksi `gorm:"foreignKey:BarangID"`
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

func (bm *BarangModel) FindBarang(pegawaiID uint) ([]Barang, error) {
	var barangs []Barang
	err := bm.DB.Where("pegawai_id = ?", pegawaiID).Find(&barangs).Error
	if err != nil {
		return nil, err
	}
	return barangs, nil
}

func (bm *BarangModel) CheckBarang(barangID uint, jumlahBarang int) (*Barang, error) {
	barang := &Barang{}
	if err := bm.DB.Where("ID = ?", barangID).First(barang).Error; err != nil {
		return nil, err
	}
	barang.Stok = barang.Stok - jumlahBarang
	barang.UpdatedAt = time.Now()

	if err := bm.DB.Save(barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}
