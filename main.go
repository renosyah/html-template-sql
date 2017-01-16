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

	rows, _ := db.Query("select * from mhs order by nim asc")

	var nim, nama, status string
	no := 1

	type mahasiswa []maba
	var data_mhs mahasiswa

	for rows.Next() {

		rows.Scan(&nim, &nama, &status)
		data := maba{
			Number: no,
			Id:     nim,
			Name:   nama,
			Text:   status,
		}
		no++
		data_mhs = append(data_mhs, data)
	}
	t, _ := template.ParseFiles("tabel.html")

	t.Execute(res, data_mhs)
}

func isi_data(res http.ResponseWriter, req *http.Request) {
	db := connect()
	defer db.Close()

	nim := req.FormValue("nim")
	nama := req.FormValue("nama")
	status := req.FormValue("status")

	_, err := db.Exec("insert into mhs value (?,?,?)", nim, nama, status)
	if err != nil {
		http.Redirect(res, req, "/", 301)
		fmt.Println("gagal input")
	}
	http.Redirect(res, req, "/", 301)
}
func hapus(res http.ResponseWriter, req *http.Request) {

	db := connect()
	defer db.Close()

	hapus := req.FormValue("hapus")

	_, err := db.Exec("delete from mhs where nim=?", hapus)
	if err != nil {
		http.Redirect(res, req, "/", 301)
		fmt.Println("gagal input")
	}
	http.Redirect(res, req, "/", 301)
}
func main() {
	http.HandleFunc("/", tampil)
	http.HandleFunc("/isi_data", isi_data)
	http.HandleFunc("/hapus", hapus)
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
	open.RunWith("http://localhost:8080/", "opera")
	http.ListenAndServe(":8080", nil)
}
