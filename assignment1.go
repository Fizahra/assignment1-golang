package main

import (
	"fmt"
	"os"
)

type absen struct {
	nabsen    int
	nama      string
	alamat    string
	pekerjaan string
	motivasi  string
}

func main() {
	var a1 = absen{1, "Fitri A.R", "Bandung", "Mahasiswa", "Ingin mempelajari bahasa pemrograman baru"}
	var a2 = absen{2, "Surya Dio A.W", "Malang", "Mahasiswa", "Ingin menjadi Back End Programmer"}
	var a3 = absen{3, "Dinar Apriliana", "Subang", "Mahasiswa", "Karena terlihat menarik"}
	var a4 = absen{4, "Luthfia Handayani", "Bandung", "Mahasiswa", "Ingin menguasai banyak bahasa pemrograman"}
	var a5 = absen{5, "Nabila Febiyanti", "Bandung", "Mahasiswa", "Karena terlihat mudah dan banyak digunakan"}

	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "1" {
			fmt.Println("Absen no.", a1.nabsen, "\nNama :", a1.nama, "\nAlamat :", a1.alamat,
				"\nPekerjaan :", a1.pekerjaan, "\nAlasan belajar Golang :", a1.motivasi)
		} else if command == "2" {
			fmt.Println("Absen no.", a2.nabsen, "\nNama :", a2.nama, "\nAlamat :", a2.alamat,
				"\nPekerjaan :", a2.pekerjaan, "\nAlasan belajar Golang :", a2.motivasi)
		} else if command == "3" {
			fmt.Println("Absen no.", a3.nabsen, "\nNama :", a3.nama, "\nAlamat :", a3.alamat,
				"\nPekerjaan :", a3.pekerjaan, "\nAlasan belajar Golang :", a3.motivasi)
		} else if command == "4" {
			fmt.Println("Absen no.", a4.nabsen, "\nNama :", a4.nama, "\nAlamat :", a4.alamat,
				"\nPekerjaan :", a4.pekerjaan, "\nAlasan belajar Golang :", a4.motivasi)
		} else if command == "5" {
			fmt.Println("Absen no.", a5.nabsen, "\nNama :", a5.nama, "\nAlamat :", a5.alamat,
				"\nPekerjaan :", a5.pekerjaan, "\nAlasan belajar Golang :", a5.motivasi)
		} else if command == "a" {
			fmt.Println("Daftar Absen Kelas Golang")
			fmt.Println("No Absen - Nama - Alamat - Pekerjaan - Alasan belajar Golang")
			fmt.Println(a1)
			fmt.Println(a2)
			fmt.Println(a3)
			fmt.Println(a4)
			fmt.Println(a5)
		}
	}
	ketikNomor()

}

func ketikNomor() {
	fmt.Println("\nSilakan masukan angka atau ketik a untuk melihat daftar absen disebelah command run")
}
