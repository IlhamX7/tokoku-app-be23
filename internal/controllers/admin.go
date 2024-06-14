package controllers

import (
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
