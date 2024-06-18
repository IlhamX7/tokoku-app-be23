package main

import (
	"errors"
	"fmt"
	"tokoku-app-be23/configs"
	"tokoku-app-be23/internal/controllers"
	"tokoku-app-be23/internal/models"

	"gorm.io/gorm"
)

func TestConnect() (*gorm.DB, error) {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return nil, err
	}
	if err := connection.AutoMigrate(&models.Admin{}); err != nil {
		return nil, err
	}
	if err := connection.First(&models.Admin{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		admin := models.Admin{Username: "admin", Password: "admin"}

		am := models.NewAdminModel(connection)
		ac := controllers.NewAdminController(am)

		_, err := ac.CreateAdmin(admin)
		if err != nil {
			fmt.Println("error ketika menambahkan admin pertama")
			return nil, err
		}
		fmt.Println("berhasil menambahkan admin pertama")
	}
	return connection, nil
}

func App() {
	connection, _ := TestConnect()

	am := models.NewAdminModel(connection)
	ac := controllers.NewAdminController(am)

	var inputMenu int
	var menus []int
	for inputMenu != 9 {
		fmt.Println("")
		fmt.Println("Selamat datang di Tokoku app")
		fmt.Println("Silahkan pilih menu yang tersedia:")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login Pegawai")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

		menus = []int{1, 2, 9}
		for _, val := range menus {
			if val == inputMenu {
				break
			}
			fmt.Println("Maaf menu yang anda masukan tidak ada, silahkan pilih lagi")
			break
		}
		if inputMenu == 1 {
			var isLogin = true
			var inputMenu2 int
			data, err := ac.LoginAdmin()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}
			for isLogin {
				fmt.Println("")
				fmt.Println("Selamat datang ", data.Username, ",")
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah pegawai")
				fmt.Println("2. Update pegawai")
				fmt.Println("3. Hapus pegawai")
				fmt.Println("4. Tampilkan daftar pegawai")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				} else if inputMenu2 == 2 {
					fmt.Println("masuk 2")
				}
			}
		}
	}
	fmt.Println("Terima kasih")
}

func main() {
	App()
}
