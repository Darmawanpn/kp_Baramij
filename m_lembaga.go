package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)

type Lembaga struct {
	id              string `json:"id"`
	XKodeSekolah 	string `json:"XKodeSekolah"`
	Yayasan 		string `json:"Yayasan"`
	Nama 			string `json:"Nama"`
	Jurusan 		string `json:"Jurusan"`
	Alamat 			string `json:"Alamat"`
	Ketua         	string `json:"Ketua"`
	Puket1      	string `json:"Puket1"`
	Puket2        	string `json:"Puket2"`
	Puket3 			string `json:"Puket3"`
	Email   		string `json:"Email"`
	Provinsi 		string `json:"Provinsi"`
	Kab_kota  		string `json:"Kab_kota"`
	Kecamatan 		string `json:"Kecamatan"`
	Kelurahan 		string `json:"Kelurahan"`
	Kodepos         string `json:"Kodepos"`
}


// Get all orders M_lembaga

func getLembaga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m_lembaga []Lembaga

	sql := `SELECT
	            id,
				IFNULL(XKodeSekolah,''),
				IFNULL(Yayasan,''),
				IFNULL(Nama,'') Nama,
				IFNULL(Jurusan,'') Jurusan,
				IFNULL(Alamat,'') Alamat,
				IFNULL(Ketua,'') Ketua,
				IFNULL(Puket1,'') Puket1,
				IFNULL(Puket2,'') Puket2 ,
				IFNULL(Puket3,'') Puket3,
				IFNULL(Email,'') Email,
				IFNULL(Provinsi,'') Provinsi,
				IFNULL(Kab_kota,'') Kab_kota,
				IFNULL(Kecamatan,'') Kecamatan ,
				IFNULL(Kelurahan,'') Kelurahan,
				IFNULL(Kodepos,'') Kodepos
			FROM m_lembaga`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var lembaga Lembaga
		err := result.Scan(&lembaga.id, &lembaga.XKodeSekolah, &lembaga.Yayasan, &lembaga.Nama, &lembaga.Jurusan, &lembaga.Alamat,
			&lembaga.Ketua, &lembaga.Puket1, &lembaga.Puket2, &lembaga.Puket3, &lembaga.Email, &lembaga.Provinsi,
			&lembaga.Kab_kota, &lembaga.Kecamatan, &lembaga.Kelurahan, &lembaga.Kodepos)

		if err != nil {
			panic(err.Error())
		}
		m_lembaga = append(m_lembaga, lembaga)
	}

	json.NewEncoder(w).Encode(m_lembaga)
}

func createLembaga(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		XKodeSekolah := r.FormValue("xKodeSekolah")
		Yayasan := r.FormValue("yayasan")
		Nama := r.FormValue("nama")
		Jurusan := r.FormValue("jurusan")
		Alamat := r.FormValue("alamat")
		Ketua := r.FormValue("ketua")
		Puket1 := r.FormValue("puket1")
		Puket2 := r.FormValue("puket2")
		Puket3 := r.FormValue("puket3")
		Email := r.FormValue("email")
		Provinsi := r.FormValue("provinsi")
		Kab_kota := r.FormValue("kab_kota")
		Kecamatan := r.FormValue("kecamatan")
		Kelurahan := r.FormValue("kelurahan")
		Kodepos := r.FormValue("kodepos")
		

		stmt, err := db.Prepare("INSERT INTO customers (id,XKodeSekolah,Yayasan,Nama,Jurusan,Alamat,Ketua,Puket1,Puket2,Puket3,Email,Provinsi,Kab_kota,Kecamatan,Kelurahan,Kodepos) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(id,XKodeSekolah,Yayasan,Nama,Jurusan,Alamat,Ketua,Puket1,Puket2,Puket3,Email,Provinsi,Kab_kota,Kecamatan,Kelurahan,Kodepos)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}
	}
}

func updateLembaga(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newYayasan := r.FormValue("Yayasan")

		stmt, err := db.Prepare("UPDATE m_lembaga SET Yayasan = ? WHERE ID = ?")

		_, err = stmt.Exec(newYayasan, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Lembaga with ID = %s was updated", params["id"])
	}
}

func deleteLembaga(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM m_lembaga WHERE ID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Lembaga with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var m_lembaga []Lembaga

	id := r.FormValue("id")
	XKodeSekolah := r.FormValue("xKodeSekolah")

	sql := `SELECT
	            id,
				IFNULL(XKodeSekolah,''),
				IFNULL(Yayasan,''),
				IFNULL(Nama,'') Nama,
				IFNULL(Jurusan,'') Jurusan,
				IFNULL(Alamat,'') Alamat,
				IFNULL(Ketua,'') Ketua,
				IFNULL(Puket1,'') Puket1,
				IFNULL(Puket2,'') Puket2 ,
				IFNULL(Puket3,'') Puket3,
				IFNULL(Email,'') Email,
				IFNULL(Provinsi,'') Provinsi,
				IFNULL(Kab_kota,'') Kab_kota,
				IFNULL(Kecamatan,'') Kecamatan ,
				IFNULL(Kelurahan,'') Kelurahan,
				IFNULL(Kodepos,'') Kodepos
			FROM m_lembaga WHERE id = ? AND XKodeSekolah = ?`

	result, err := db.Query(sql, id, XKodeSekolah)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var lembaga Lembaga

	for result.Next() {

		err := result.Scan(&lembaga.id, &lembaga.XKodeSekolah, &lembaga.Yayasan, &lembaga.Nama, &lembaga.Jurusan, &lembaga.Alamat,
			&lembaga.Ketua, &lembaga.Puket1, &lembaga.Puket2, &lembaga.Puket3, &lembaga.Email, &lembaga.Provinsi,
			&lembaga.Kab_kota, &lembaga.Kecamatan, &lembaga.Kelurahan, &lembaga.Kodepos)

		if err != nil {
			panic(err.Error())
		}

		m_lembaga = append(m_lembaga, lembaga)
	}

	json.NewEncoder(w).Encode(m_lembaga)

}


// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/m_lembaga/{id}", getLembaga).Methods("GET")
	r.HandleFunc("/m_lembaga", createLembaga).Methods("POST")
	r.HandleFunc("/m_lembaga/{id}", updateLembaga).Methods("PUT")
	r.HandleFunc("/m_lembaga/{id}", deleteLembaga).Methods("DELETE")

	//new
	r.HandleFunc("/getLembaga", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8383", r))
}
