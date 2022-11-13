package main

import (
	config "fajars/rowsql/config"
	"fajars/rowsql/controllers"
	entity "fajars/rowsql/entity"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.KonekDB()

	defer db.Close() // menutup koneksi

	//buat mekanisme menu
	fmt.Println("MENU:\n1. Baca data\n2. Tambah data\n3. update")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			dataUser, errGetAll := controllers.Userdata(db) // memanggil fungsi di controllers
			if errGetAll != nil {                           // cek apakah ada error dari getAllUser
				fmt.Println("error get all users")
			}
			// fmt.Println("data all", dataUser)
			for _, value := range dataUser { // membaca/menampilkan  seluruh data user yang telah ditampung di variable slice
				fmt.Printf("id: %s, nama: %s, email: %s, phone: %s\n", value.Id, value.Nama, value.Gender, value.Status)
			}
		}
	case 2:
		{

			newUser := entity.User{}

			fmt.Println("masukkan nama user")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("masukkan gender user")
			fmt.Scanln(&newUser.Gender)
			fmt.Println("masukkan status user")
			fmt.Scanln(&newUser.Status)

			controllers.InsertData(db, newUser)

		}
	case 3:
		{
			update := entity.User{}
			fmt.Println("update id yang akan diubah")
			fmt.Scanln(&update.Id)
			fmt.Println("Update nama User")
			fmt.Scanln(&update.Nama)
			fmt.Println("Update gender User")
			fmt.Scanln(&update.Gender)
			fmt.Println("Update status User")
			fmt.Scanln(&update.Status)

			controllers.Update(db, update)
		}

	case 4:
		{
			fmt.Println("delete")
		}

	case 5:
		{
			fmt.Println("baca data by id")
		}

	}

}
