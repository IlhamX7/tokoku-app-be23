package controllers

import (
	"time"
	"tokoku-app-be23/internal/models"

	"github.com/google/uuid"
)

type NotaTransaksiController struct {
	model *models.NotaTransaksiModel
}

func NewNotaTransaksiController(m *models.NotaTransaksiModel) *NotaTransaksiController {
	return &NotaTransaksiController{
		model: m,
	}
}

func (ntc *NotaTransaksiController) CreateNotaTransaksi(pegawaiId uint, customerId uint) (bool, error) {
	var newData models.NotaTransaksi

	newData.PegawaiId = pegawaiId
	newData.CustomerId = customerId

	newUUID := uuid.New()
	uuidString := newUUID.String()
	newData.KodeTransaksi = uuidString

	newData.NamaTransaksi = InputString("Masukkan nama transaksi ")
	newData.TanggalTransaksi = time.Now()

	newData.TotalBarang = 0
	newData.Keterangan = InputString("Masukan keterangan ")

	_, err := ntc.model.CreateNotaTransaksi(newData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ntc *NotaTransaksiController) DeleteNotaTransaksi(id uint) (bool, error) {
	_, err := ntc.model.DeleteNotaTransaksi(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
