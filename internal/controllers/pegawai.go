package controllers

import (
	"errors"
	"tokoku-app-be23/internal/models"
)

type PegawaiController struct {
	model *models.PegawaiModel
}

func NewPegawaiController(m *models.PegawaiModel) *PegawaiController {
	return &PegawaiController{
		model: m,
	}
}

func (pc *PegawaiController) AddPegawai(id uint) (bool, error) {
	var newData models.Pegawai

	newData.Username = InputString("Masukkan username ")
	newData.Password = InputString("Masukkan password ")
	newData.Email = InputString("Masukkan email ")
	newData.AdminId = id

	_, err := pc.model.AddPegawai(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (pc *PegawaiController) UpdatePegawai(id uint) (bool, error) {
	var updateData models.Pegawai

	updateData.Username = InputString("Masukkan username baru ")
	updateData.Password = InputString("Masukkan password baru ")
	updateData.Email = InputString("Masukkan email baru ")

	_, err := pc.model.UpdatePegawai(id, updateData.Username, updateData.Password, updateData.Email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (pc *PegawaiController) DeletePegawai(id uint) (bool, error) {
	_, err := pc.model.DeletePegawai(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (pc *PegawaiController) FindPegawai(id uint) ([]models.ResponsePegawai, error) {
	data, err := pc.model.FindPegawai(id)
	if err != nil {
		return nil, err
	}

	responsePegawai := make([]models.ResponsePegawai, len(*data))
	for i, val := range *data {
		responsePegawai[i] = models.ResponsePegawai{
			ID:       val.ID,
			Username: val.Username,
			Password: val.Password,
			Email:    val.Email,
		}
	}

	return responsePegawai, nil
}

func (pc *PegawaiController) LoginPegawai() (models.Pegawai, error) {
	username := InputString("Masukkan username ")
	password := InputString("Masukkan password ")
	result, err := pc.model.LoginPegawai(username, password)
	if err != nil {
		return models.Pegawai{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}
