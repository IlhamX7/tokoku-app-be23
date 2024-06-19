package controllers

import (
	"fmt"
	"strconv"
	"tokoku-app-be23/internal/models"
)

type BarangController struct {
	model *models.BarangModel
}

func NewBarangController(m *models.BarangModel) *BarangController {
	return &BarangController{
		model: m,
	}
}

// Fungsi untuk mengambil input string
func InputData(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}

// Fungsi untuk mengambil input integer
func InputInt(prompt string) (int, error) {
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

func (bc *BarangController) AddBarang(pegawaiId uint) (bool, error) {
	var newBarang models.Barang
	var err error

	newBarang.KodeBarang = InputData("Masukkan kode barang")
	newBarang.NamaBarang = InputData("Masukkan nama barang")

	newBarang.Stok, err = InputInt("Masukkan stok barang")
	if err != nil {
		return false, err
	}

	newBarang.Harga, err = InputInt("Masukkan harga barang")
	if err != nil {
		return false, err
	}

	newBarang.Keterangan = InputData("Masukkan keterangan barang")
	newBarang.PegawaiId = pegawaiId

	_, err = bc.model.AddBarang(newBarang)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bc *BarangController) UpdateBarang(id uint) (bool, error) {
	var updateData models.Barang
	var err error

	updateData.KodeBarang = InputData("Masukkan kode barang baru")
	updateData.NamaBarang = InputData("Masukkan nama barang baru")

	updateData.Stok, err = InputInt("Masukkan stok barang baru")
	if err != nil {
		return false, err
	}

	updateData.Harga, err = InputInt("Masukkan harga barang baru")
	if err != nil {
		return false, err
	}

	updateData.Keterangan = InputData("Masukkan keterangan barang baru")

	_, err = bc.model.UpdateBarang(id, updateData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bc *BarangController) DeleteBarang(id uint) (bool, error) {
	_, err := bc.model.DeleteBarang(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bc *BarangController) FindBarang(id uint) ([]models.ResponseBarang, error) {
	data, err := bc.model.FindBarang(id)
	if err != nil {
		return nil, err
	}

	responseBarang := make([]models.ResponseBarang, len(*data))
	for i, val := range *data {
		responseBarang[i] = models.ResponseBarang{
			ID:         val.ID,
			KodeBarang: val.KodeBarang,
			NamaBarang: val.NamaBarang,
			Stok:       val.Stok,
			Harga:      val.Harga,
			Keterangan: val.Keterangan,
		}
	}

	return responseBarang, nil
}

func (bc *BarangController) GetBarang(id uint) (bool, error) {
	_, err := bc.model.GetBarang(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
