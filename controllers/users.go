package controllers

import (
	"database/sql"
	entity "fajars/rowsql/entity"
	"fmt"
	"log"
)

func Userdata(db *sql.DB) ([]entity.User, error) {
	result, errselect := db.Query("SELECT id,name,gender,status FROM users") // proses menjalankana query SQL

	if errselect != nil { //handling error saat proses menjalankan query
		return nil, errselect
	}
	var datausers []entity.User

	for result.Next() {
		var rowuser entity.User
		errscan := result.Scan(&rowuser.Id, &rowuser.Nama, &rowuser.Gender, &rowuser.Status) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumny
		if errscan != nil {                                                                  // handling ketika ada error pada saat proses scanning
			// log.Fatal("eror scan", errscan.Error())
			return nil, errscan
		}
		// fmt.Printf("id: %s, nama: %s, gender:%s, status%s\n", rowuser.id, rowuser.nama, rowuser.gender, rowuser.status) // menampilkan data hasil pembacaan dari db
		datausers = append(datausers, rowuser)
	}
	return datausers, nil
}

func InsertData(db *sql.DB, newUser entity.User) (sql.Result, error) {

	var query = "INSERT INTO users(name,gender,status) VALUES (?,?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(newUser.Nama, newUser.Gender, newUser.Status)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}

func Update(db *sql.DB, update entity.User) (sql.Result, error) {
	var query = "UPDATE users set name = ?, gender = ?, status = ? where id = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare update", errPrepare.Error())
	}
	result, errExec := statement.Exec(update.Nama, update.Gender, update.Status, update.Id)

	if errExec != nil {
		log.Fatal("erorr Exec update", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}
