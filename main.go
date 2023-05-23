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

type Forum struct {
	pertanyaan PertanyaanArr
}

type Pertanyaan struct {
	pasien  Pasien
	judul   string
	topik   TopikArr
	replies ReplyArr
}

type Topik string

type TopikArr struct {
	info [ARR_STATIC_MAX]Topik
	n    int
}

type PertanyaanArr struct {
	info [ARR_STATIC_MAX]Pertanyaan
	n    int
}

type Reply struct {
	nama, message, tipe string
}

type ReplyArr struct {
	info [ARR_STATIC_MAX]Reply
	n    int
}

func Menu() {
	fmt.Println(`
Konsultasi Kesehatan
--------------------
1. Daftar
2. Login
3. Logout
4. Forum
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

func PasienPush(arr *PasienArr, x Pasien) {
	if arr.n < ARR_STATIC_MAX {
		arr.info[arr.n] = x
		arr.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Pasien")
	}
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
		PasienPush(arr, pasien)
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
		arr[i] = key
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

/* Forum F() +
 */

func ForumPush(f *Forum, p Pertanyaan) {
	if f.pertanyaan.n < ARR_STATIC_MAX {
		f.pertanyaan.info[f.pertanyaan.n] = p
		f.pertanyaan.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Forum")
	}
}

func ForumPrint(f Forum) {
	for i := 0; i < f.pertanyaan.n; i++ {
		fmt.Print(i + 1)
		fmt.Print(" Judul:\n", f.pertanyaan.info[i].judul, "\n")

		fmt.Println("Diskusi:")
		for j := 0; j < f.pertanyaan.info[i].replies.n; j++ {
			fmt.Print("[", f.pertanyaan.info[i].replies.info[j].nama, "(", f.pertanyaan.info[i].replies.info[j].tipe, ")]: ", f.pertanyaan.info[i].replies.info[j].message, "\n")
		}
	}
}

/* Forum F() -
 */

func PertanyaanSort__(f *Forum) {
	for i := 0; i < f.pertanyaan.n; i++ {
		for j := i + 1; j < f.pertanyaan.n; j++ {
			if f.pertanyaan.info[j].replies.n < f.pertanyaan.info[i].replies.n {
				tmp := f.pertanyaan.info[i]
				f.pertanyaan.info[i] = f.pertanyaan.info[j]
				f.pertanyaan.info[j] = tmp
			}
		}
	}
}

func PertanyaanSortAsc(f *Forum) {
	var i, j, min_idx int
	for i = 0; i < f.pertanyaan.n-1; i++ {
		min_idx = i
		for j = i + 1; j < f.pertanyaan.n; j++ {
			if f.pertanyaan.info[j].replies.n < f.pertanyaan.info[min_idx].replies.n {
				min_idx = j
			}
		}
		var temp = f.pertanyaan.info[i]
		f.pertanyaan.info[i] = f.pertanyaan.info[min_idx]
		f.pertanyaan.info[min_idx] = temp
	}
}

func PertanyaanSortDesc(f *Forum) {
	for i := 1; i < f.pertanyaan.n; i++ {
		key := f.pertanyaan.info[i]
		j := i - 1

		for j > 0 && f.pertanyaan.info[j].replies.n > key.replies.n {
			f.pertanyaan.info[j] = f.pertanyaan.info[j-1]
			j--
		}
		f.pertanyaan.info[i] = key
	}
}

func Daftar() {
	PasienDaftar(&db.pasien)
}

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
				db.user.tipe = USER_TIPE[pilihan-1]
				db.user.pasien = &db.pasien.info[idx]
				fmt.Println("[info]: Selamat datang pasien", nama)
			} else {
				fmt.Println("[info]: Login gagal")
			}
		} else if pilihan == 2 {
			idx := DokterFind(db.dokter, Dokter{nama: nama, password: password})
			if idx > -1 {
				db.user.tipe = USER_TIPE[pilihan-1]
				db.user.dokter = &db.dokter.info[idx]
				fmt.Println("[info]: Selamat datang dokter", nama)
			} else {
				fmt.Println("[info]: Login gagal")
			}
		}
	}
}

func Forum__() {
	var pilihan int

	fmt.Println("\nForum")
	fmt.Println("=====")
	fmt.Println("1. Lihat \t 2. Tambah \t 3. Reply ")

	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan >= 1 && pilihan <= 3 {
		if pilihan == 1 {
			if db.user.tipe == "DOKTER" {
				var choose string = "N"
				fmt.Print("Urutkan pertanyaan atau tidak? (Y/N): ")
				fmt.Scan(&choose)
				if choose == "y" || choose == "Y" {
					fmt.Print("Mengurutkan pertanyaan berdasarkan jumlah reply secara \n 1.Descending \t 2. Ascending \n")
					var input int
					fmt.Print("Masukkan pilihan: ")
					fmt.Scan(&input)
					if input == 1 {
						PertanyaanSortDesc(&db.forum)
					} else if input == 2 {
						PertanyaanSortAsc(&db.forum)
					}
				}
			}
			ForumPrint(db.forum)
		} else if pilihan == 2 {
			if db.user.tipe == "PASIEN" {
				var pertanyaan string
				fmt.Print("Masukan pertanyaan [masukan \"STOP\" diakhir judul untuk stop]: ")
				ScanString(&pertanyaan, "STOP")
				ForumPush(&db.forum, Pertanyaan{pasien: *db.user.pasien, judul: pertanyaan})
			} else {
				fmt.Println("[info]: Harap login sebagai pasien terlebih dahulu")
			}
		} else if pilihan == 3 {
			var reply Reply
			var i int
			if db.user.tipe == "PASIEN" || db.user.tipe == "DOKTER" {
				ForumPrint(db.forum)
				fmt.Print("Masukkan forum ke-")
				fmt.Scan(&i)
				if i > 0 || i <= db.forum.pertanyaan.n {
					if db.user.tipe == "PASIEN" {
						fmt.Print("Masukkan balasan anda: ")
						ScanString(&reply.message, "STOP")
						reply.nama = db.user.pasien.nama
						reply.tipe = "Pasien"
					} else if db.user.tipe == "DOKTER" {
						fmt.Print("Masukkan balasan anda: ")
						ScanString(&reply.message, "STOP")
						reply.nama = db.user.dokter.nama
						reply.tipe = "Dokter"
					}
					ReplyPush(&db.forum, i-1, reply)
				} else {
					fmt.Print("[info]: Pertanyaan tidak ditemukan")
				}
			} else {
				fmt.Print("[info]: Harap login terlebih dahulu")
			}

		}
	}
}

func printLogin() {
	if db.user.tipe == "PASIEN" {
		fmt.Println(db.user.pasien)
	} else {
		fmt.Println(db.user.dokter)
	}
}

// Relpy +

func ReplyPush(p *Forum, i int, x Reply) {
	if p.pertanyaan.info[i].replies.n < ARR_STATIC_MAX {
		p.pertanyaan.info[i].replies.info[p.pertanyaan.info[i].replies.n] = x
		p.pertanyaan.info[i].replies.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Balasan")
	}
}

var db Database

func main() {
	DokterPush(&db.dokter, Dokter{nama: "Helmi", password: "admin", umur: 19})
	DokterPush(&db.dokter, Dokter{nama: "Fattan", password: "admin", umur: 19})

	PasienPush(&db.pasien, Pasien{nama: "nala", password: "nala", umur: 19})
	PasienPush(&db.pasien, Pasien{nama: "aku", password: "123", umur: 19})

	db.user.tipe = USER_TIPE[0]
	db.user.pasien = &db.pasien.info[PasienFind(db.pasien, Pasien{nama: "nala", password: "nala", umur: 19})]

	ForumPush(&db.forum, Pertanyaan{pasien: *db.user.pasien, judul: "apa itu lambung"})
	ReplyPush(&db.forum, 0, Reply{nama: "nala", message: "xxx", tipe: "Pasien"})
	ReplyPush(&db.forum, 0, Reply{nama: "aku", message: "xxx", tipe: "Pasien"})
	ReplyPush(&db.forum, 0, Reply{nama: "dia", message: "xxx", tipe: "Pasien"})

	ForumPush(&db.forum, Pertanyaan{pasien: *db.user.pasien, judul: "apa itu kucing"})
	ReplyPush(&db.forum, 1, Reply{nama: "joko", message: "xxx", tipe: "Pasien"})

	PertanyaanSortDesc(&db.forum)

	for i := -1; i != 0; {
		Menu()
		fmt.Print("Masukkan: ")
		fmt.Scan(&i)

		if i == 1 {
			Daftar()
		} else if i == 2 {
			Login()
		} else if i == 3 {
			Logout()
		} else if i == 4 {
			Forum__()
		} else if i == 9 {
			printLogin()
		}
	}
}

type Database struct {
	user   User
	pasien PasienArr
	dokter DokterArr
	forum  Forum
}

var USER_TIPE = [2]string{"PASIEN", "DOKTER"}

type User struct {
	tipe   string
	pasien *Pasien
	dokter *Dokter
}

func ScanString(buf *string, stop string) {
	var str string
	fmt.Scan(&str)
	for str != stop {
		*buf += (str + " ")
		fmt.Scan(&str)
	}
}

func Logout() {
	db.user.tipe = "LOGOUT"
}
