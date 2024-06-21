package controllers

import (
	"tokoku-app-be23/internal/models"
)

type CustomerController struct {
	model *models.CustomerModel
}

func NewCustomerController(m *models.CustomerModel) *CustomerController {
	return &CustomerController{
		model: m,
	}
}

func (cc *CustomerController) AddCustomer(pegawaiId uint) (bool, error) {
	var newCustomer models.Customer

	newCustomer.NamaCustomer = InputString("Masukkan nama customer ")
	newCustomer.Address = InputString("Masukkan alamat customer ")
	newCustomer.Phone = InputString("Masukkan nomor telepon customer ")
	newCustomer.Email = InputString("Masukkan email customer ")
	newCustomer.PegawaiId = pegawaiId

	_, err := cc.model.AddCustomer(newCustomer)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cc *CustomerController) UpdateCustomer(id uint) (bool, error) {
	var updateData models.Customer

	updateData.NamaCustomer = InputString("Masukkan nama customer baru ")
	updateData.Address = InputString("Masukkan alamat customer baru ")
	updateData.Phone = InputString("Masukkan nomor telepon customer baru ")
	updateData.Email = InputString("Masukkan email customer baru ")

	_, err := cc.model.UpdateCustomer(id, updateData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cc *CustomerController) DeleteCustomer(id uint) (bool, error) {
	_, err := cc.model.DeleteCustomer(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cc *CustomerController) FindCustomer(id uint) ([]models.Customer, error) {
	data, err := cc.model.FindCustomer(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cc *CustomerController) GetCustomer(id uint) (bool, error) {
	_, err := cc.model.GetCustomer(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
