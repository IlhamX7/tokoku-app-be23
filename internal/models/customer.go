package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	PegawaiId    uint
	NamaCustomer string
	Address      string
	Phone        string
	Email        string
  NotaTransaksi []NotaTransaksi `gorm:"foreignKey:CustomerId"`
}

type ResponseCustomer struct {
	ID           uint
	NamaCustomer string
	Address      string
	Phone        string
	Email        string
}

type CustomerModel struct {
	DB *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *CustomerModel {
	return &CustomerModel{
		DB: db,
	}
}

func (cm *CustomerModel) AddCustomer(customer Customer) (*Customer, error) {
	if err := cm.DB.Create(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cm *CustomerModel) UpdateCustomer(id uint, updatedData Customer) (*Customer, error) {
	customer := &Customer{}
	if err := cm.DB.First(customer, id).Error; err != nil {
		return nil, err
	}

	customer.NamaCustomer = updatedData.NamaCustomer
	customer.Address = updatedData.Address
	customer.Phone = updatedData.Phone
	customer.Email = updatedData.Email
	customer.UpdatedAt = time.Now()

	if err := cm.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (cm *CustomerModel) DeleteCustomer(id uint) (*Customer, error) {
	customer := &Customer{}
	if err := cm.DB.First(customer, id).Error; err != nil {
		return nil, err
	}

	if err := cm.DB.Delete(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (cm *CustomerModel) FindCustomer(pegawaiID uint) (*[]Customer, error) {
	customers := &[]Customer{}
	if err := cm.DB.Where("pegawai_id = ?", pegawaiID).Find(customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (cm *CustomerModel) GetCustomer(customerID uint) (*Customer, error) {
	customer := &Customer{}
	if err := cm.DB.Where("ID = ?", customerID).First(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
