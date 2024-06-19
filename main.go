package main

import (
	"errors"
	"fmt"
	"strconv"
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
	if err := connection.AutoMigrate(&models.Pegawai{}); err != nil {
		return nil, err
	}
	if err := connection.AutoMigrate(&models.Customer{}); err != nil {
		return nil, err
	}
	if err := connection.AutoMigrate(&models.Barang{}); err != nil {
		return nil, err
	}
	if err := connection.AutoMigrate(&models.NotaTransaksi{}); err != nil {
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

func App() {
	connection, _ := TestConnect()

	am := models.NewAdminModel(connection)
	ac := controllers.NewAdminController(am)

	pu := models.NewPegawaiModel(connection)
	pc := controllers.NewPegawaiController(pu)

	cm := models.NewCustomerModel(connection)
	cc := controllers.NewCustomerController(cm)

	bm := models.NewBarangModel(connection)
	bc := controllers.NewBarangController(bm)

	ntm := models.NewNotaTransaksiModel(connection)
	ntc := controllers.NewNotaTransaksiController(ntm)

	var inputMenu int
	for inputMenu != 9 {
		fmt.Println("")
		fmt.Println("Selamat datang di Tokoku app")
		fmt.Println("Silahkan pilih menu yang tersedia:")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login Pegawai")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

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
				fmt.Println("Selamat datang", data.Username)
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah pegawai")
				fmt.Println("2. Update pegawai")
				fmt.Println("3. Hapus pegawai")
				fmt.Println("4. Tampilkan daftar pegawai")
				fmt.Println("5. Hapus customer")
				fmt.Println("6. Hapus barang")
				fmt.Println("7. Hapus nota transaksi")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				} else if inputMenu2 == 1 {
					_, err := pc.AddPegawai(data.ID)
					if err != nil {
						fmt.Println("error ketika menambahkan pegawai")
						return
					}
					fmt.Println("berhasil menambahkan pegawai")
				} else if inputMenu2 == 2 {
					var id uint
					fmt.Print("Masukkan id pegawai ")
					fmt.Scanln(&id)
					_, err := pc.UpdatePegawai(id)
					if err != nil {
						fmt.Println("error ketika mengubah pegawai")
						return
					}
					fmt.Println("berhasil mengubah pegawai")
				} else if inputMenu2 == 3 {
					var id uint
					fmt.Print("Masukkan id pegawai ")
					fmt.Scanln(&id)
					_, err := pc.DeletePegawai(id)
					if err != nil {
						fmt.Println("error ketika menghapus pegawai")
						return
					}
					fmt.Println("berhasil menghapus pegawai")
				} else if inputMenu2 == 4 {
					data, err := pc.FindPegawai(data.ID)
					if err != nil {
						fmt.Println("error ketika menampilkan daftar pegawai")
						return
					}
					fmt.Println("berhasil menampilkan daftar pegawai")
					for i, pegawai := range data {
						fmt.Printf("Pegawai %d:\nId: %d\nUsername: %s\nEmail: %s\n", i+1, pegawai.ID, pegawai.Username, pegawai.Email)
					}
				} else if inputMenu2 == 5 {
					var id uint
					fmt.Print("Masukkan id customer ")
					fmt.Scanln(&id)
					_, err := cc.DeleteCustomer(id)
					if err != nil {
						fmt.Println("error ketika menghapus customer")
						return
					}
					fmt.Println("berhasil menghapus customer")
				} else if inputMenu2 == 6 {
					var id uint
					fmt.Print("Masukkan id barang ")
					fmt.Scanln(&id)
					_, err := bc.DeleteBarang(id)
					if err != nil {
						fmt.Println("error ketika menghapus barang")
						return
					}
					fmt.Println("berhasil menghapus barang")
				} else if inputMenu2 == 7 {
					var id uint
					fmt.Print("Masukkan id nota transaksi ")
					fmt.Scanln(&id)
					_, err := ntc.DeleteNotaTransaksi(id)
					if err != nil {
						fmt.Println("error ketika menghapus nota transaksi")
						return
					}
					fmt.Println("berhasil menghapus nota transaksi")
				} else {
					fmt.Println("maaf menu yang anda pilih tidak ada, silahkan pilih lagi")
					continue
				}
			}
		} else if inputMenu == 2 {
			var isLogin = true
			var inputMenu2 int
			data, err := pc.LoginPegawai()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}
			for isLogin {
				fmt.Println("")
				fmt.Println("Selamat datang", data.Username)
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah barang baru")
				fmt.Println("2. Edit informasi barang")
				fmt.Println("3. Update stok barang")
				fmt.Println("4. Tampilkan daftar barang")
				fmt.Println("5. Tambah customer baru")
				fmt.Println("6. Create nota transaksi untuk customer")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				} else if inputMenu2 == 1 {
					_, err := bc.AddBarang(data.ID)
					if err != nil {
						fmt.Println("error ketika menambahkan barang")
						return
					}
					fmt.Println("berhasil menambahkan barang")
				} else if inputMenu2 == 2 {
					var id uint
					fmt.Print("Masukkan id barang ")
					fmt.Scanln(&id)
					_, err := bc.UpdateBarang(id)
					if err != nil {
						fmt.Println("error ketika mengubah barang")
						return
					}
					fmt.Println("berhasil mengubah barang")
				} else if inputMenu2 == 3 {
					var id uint
					fmt.Print("Masukkan id barang ")
					fmt.Scanln(&id)
					_, err := bc.UpdateBarang(id)
					if err != nil {
						fmt.Println("error ketika mengubah barang")
						return
					}
					fmt.Println("berhasil mengubah stok barang")
				} else if inputMenu2 == 4 {
					data, err := bc.FindBarang(data.ID)
					if err != nil {
						fmt.Println("error ketika menampilkan daftar barang")
						return
					}
					fmt.Println("berhasil menampilkan daftar barang")
					for i, barang := range data {
						fmt.Printf("Barang %d:\nId: %d\nKode: %s\nNama: %s\nHarga: %d\n\n", i+1, barang.ID, barang.KodeBarang, barang.NamaBarang, barang.Harga)
					}
				} else if inputMenu2 == 5 {
					_, err := cc.AddCustomer(data.ID)
					if err != nil {
						fmt.Println("error ketika menambahkan customer")
						return
					}
					fmt.Println("berhasil menambahkan customer")
				} else if inputMenu2 == 6 {
					customerId, err := InputUint("Silahkan pilih id customer ")
					if err != nil {
						fmt.Println("error ketika memilih id customer")
						return
					}
					_, err = cc.GetCustomer(customerId)
					if err != nil {
						fmt.Println("error ketika memilih id customer")
						return
					}
					_, err = ntc.CreateNotaTransaksi(data.ID, customerId)
					if err != nil {
						fmt.Println("error ketika membuat nota transaksi")
						return
					}
					fmt.Println("berhasil membuat nota transaksi")
				}
			}
		} else if inputMenu == 9 {
			break
		}
	}
	fmt.Println("Terima kasih")
}

func main() {
	App()
}
