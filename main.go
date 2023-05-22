package main

import (
	"fmt"
)

const ARR_STATIC_MAX int = 1024

type Pasien struct {
	nama, password string
	umur           int
}

type PasienArr struct {
	info [ARR_STATIC_MAX]Pasien
	n    int
}

type Dokter struct {
	nama, password string
	umur           int
}

type DokterArr struct {
	info [ARR_STATIC_MAX]Dokter
	n    int
}

func Menu() {
	fmt.Println(`
Konsultasi Kesehatan
--------------------
1. Daftar
2. Login
3. Forum
0. Keluar
    `)
}

/* Pasien F() +
 */
func PasienFind(arr PasienArr, x Pasien) int {
	var idx int = -1
	for i := 0; i < arr.n && idx == -1; i++ {
		if x.nama == arr.info[i].nama && x.password == arr.info[i].password {
			idx = i
		}
	}
	return idx
}

func PasienDaftar(arr *PasienArr) {
	var pasien Pasien
	var hasil int

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&pasien.nama)

	fmt.Print("Masukkan Password: ")
	fmt.Scan(&pasien.password)

	fmt.Print("Masukkan Umur: ")
	fmt.Scan(&pasien.umur)

	hasil = PasienFind(*arr, pasien)
	if hasil == -1 {
		if arr.n < ARR_STATIC_MAX {
			arr.info[arr.n] = pasien
			arr.n++

			fmt.Println("[info]: Pengguna terdaftar")
		} else {
			fmt.Println("[info]: Pengguna gagal terdaftar")
		}
	} else {
		fmt.Println("[info]: Pengguna sudah terdaftar")
	}
}

func PasienSort(arr []Pasien, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j > 0 && arr[j].umur > key.umur {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = key
	}
}

/* Pasien F() -
 */

/* Dokter F() +
 */
func DokterFind(arr DokterArr, x Dokter) int {
	var idx int = -1
	for i := 0; i < arr.n && idx == -1; i++ {
		if x.nama == arr.info[i].nama && x.password == arr.info[i].password {
			idx = i
		}
	}
	return idx
}

func DokterPush(arr *DokterArr, x Dokter) {
	if arr.n < ARR_STATIC_MAX {
		arr.info[arr.n] = x
		arr.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Dokter")
	}
}

/* Dokter F() -
 */

func Login() {
	var pilihan int

	fmt.Println("\nLogin sebagai")
	fmt.Println("=============")
	fmt.Println("1. Pasien \t 2. Dokter")

	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan >= 1 && pilihan <= 2 {
		var nama, password string

		fmt.Print("Masukan Nama: ")
		fmt.Scan(&nama)

		fmt.Print("Masukan Password: ")
		fmt.Scan(&password)

		if pilihan == 1 {
			idx := PasienFind(db.pasien, Pasien{nama: nama, password: password})
			if idx > -1 {
				fmt.Println("[info]: Selamat datang pasien", nama)
			} else {
				fmt.Println("[info]: Login gagal")
			}
		} else if pilihan == 2 {
			idx := DokterFind(db.dokter, Dokter{nama: nama, password: password})
			if idx > -1 {
				fmt.Println("[info]: Selamat datang dokter", nama)
			} else {
				fmt.Println("[info]: Login gagal")
			}
		}
	}
}

var db Database

func main() {
	DokterPush(&db.dokter, Dokter{nama: "Helmi", password: "admin"})
	DokterPush(&db.dokter, Dokter{nama: "Fattan", password: "admin"})

	for i := -1; i != 0; {
		Menu()
		fmt.Print("Masukkan: ")
		fmt.Scan(&i)

		if i == 1 {
			PasienDaftar(&db.pasien)
		} else if i == 2 {
			Login()
		}
	}
}

type Database struct {
	pasien PasienArr
	dokter DokterArr
}
