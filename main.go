package main

import (
	"fmt"
	"tokoku-app-be23/configs"
	"tokoku-app-be23/internal/controllers"
	"tokoku-app-be23/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	connection.AutoMigrate(&models.Admin{})

	am := models.NewAdminModel(connection)
	ac := controllers.NewAdminController(am)
	fmt.Println(ac)
}
