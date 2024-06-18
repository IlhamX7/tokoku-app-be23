package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"tokoku-app-be23/internal/models"
)

type AdminController struct {
	model *models.AdminModel
}

func NewAdminController(m *models.AdminModel) *AdminController {
	return &AdminController{
		model: m,
	}
}

func InputString(pertanyaan string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pertanyaan)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	return input
}

func (ac *AdminController) CreateAdmin(admin models.Admin) (bool, error) {
	_, err := ac.model.CreateAdmin(admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ac *AdminController) LoginAdmin() (models.Admin, error) {
	username := InputString("Masukkan username ")
	password := InputString("Masukkan password ")
	result, err := ac.model.LoginAdmin(username, password)
	if err != nil {
		return models.Admin{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}
