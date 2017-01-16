package main

import "html/template"
import "net/http"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/skratchdot/open-golang/open"
import "os"
import "fmt"

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "renosyah:renosyah@/go_html")
	err = db.Ping()
	if err != nil {
		fmt.Println("database tidak bisa dihubungi")
		os.Exit(0)

	}
	return db

}

type maba struct { // isi tidak boleh di deklarasikan sembarangan
	Number int
	Id     string
	Name   string
	Text   string
}

func tampil(res http.ResponseWriter, req *http.Request) {
	db := connect()
	defer db.Close()

	rows, _ := db.Query("select * from mhs")

	var nim, nama, status string
	var a int = 1

	type mahasiswa []maba
	var data_mhs mahasiswa

	for rows.Next() {

		rows.Scan(&nim, &nama, &status)
		data := maba{
			Number: a,
			Id:     nim,
			Name:   nama,
			Text:   status,
		}
		a++
		data_mhs = append(data_mhs, data)
	}
	t, _ := template.ParseFiles("tabel.html")

	t.Execute(res, data_mhs)
}

func main() {
	http.HandleFunc("/", tampil)
	open.RunWith("http://localhost:8080/", "opera")
	http.ListenAndServe(":8080", nil)
}
