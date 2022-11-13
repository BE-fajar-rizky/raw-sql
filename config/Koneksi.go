package configs

import (
	"database/sql"
	"fmt"
	"log"
)

func KonekDB() *sql.DB {
	var koneksi = "root:@tcp(127.0.0.1:3306)/altera_be13?parseTime=true"
	db, err := sql.Open("mysql", koneksi) // membuka koneksi ke daatabase

	if err != nil { // pengecekan error yang terjadi ketika proses open connectio
		log.Fatal("eror", err.Error())
	}
	errPing := db.Ping() //check terkoneksi atau ga
	if errPing != nil {  //handling error ketika gagal konek ke db
		log.Fatal("eror", errPing.Error())
	} else {
		fmt.Println("berhasil")
	}
	return db
}
