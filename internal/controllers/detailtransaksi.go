package controllers

import (
	"fmt"
	"strconv"
	"tokoku-app-be23/internal/models"
)

type DetailTransaksiController struct {
	model *models.DetailTransaksiModel
}

func NewDetailTransaksiController(m *models.DetailTransaksiModel) *DetailTransaksiController {
	return &DetailTransaksiController{
		model: m,
	}
}

func InputDt(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}

func InputUint(prompt string) (uint, error) {
	var input string
	fmt.Print(prompt + ": ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if value < 0 {
		return 0, fmt.Errorf("negative value entered: %d", value)
	}
	return uint(value), nil
}

func InputInteger(prompt string) (int, error) {
	var input string
	fmt.Print(prompt + ": ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (dtc *DetailTransaksiController) AddDetailTransaksi(notaTransaksiID uint) (bool, error) {
	var newDetail models.DetailTransaksi
	var err error

	newDetail.NotaTransaksiID = notaTransaksiID
	newDetail.BarangID, err = InputUint("Masukkan ID barang")
	if err != nil {
		return false, err
	}

	newDetail.JumlahBarang, err = InputInt("Masukkan jumlah barang")
	if err != nil {
		return false, err
	}

	newDetail.HargaBarang, err = InputInt("Masukkan harga barang")
	if err != nil {
		return false, err
	}

	_, err = dtc.model.AddDetailTransaksi(newDetail)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dtc *DetailTransaksiController) DeleteDetailTransaksi(id uint) (bool, error) {
	_, err := dtc.model.DeleteDetailTransaksi(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dtc *DetailTransaksiController) FindDetailTransaksi(notaTransaksiID uint) ([]models.ResponseDetailTransaksi, error) {
	data, err := dtc.model.FindDetailTransaksi(notaTransaksiID)
	if err != nil {
		return nil, err
	}

	responseDetail := make([]models.ResponseDetailTransaksi, len(*data))
	for i, val := range *data {
		responseDetail[i] = models.ResponseDetailTransaksi{
			ID:              val.ID,
			NotaTransaksiID: val.NotaTransaksiID,
			BarangID:        val.BarangID,
			JumlahBarang:    val.JumlahBarang,
			HargaBarang:     val.HargaBarang,
		}
	}

	return responseDetail, nil
}
