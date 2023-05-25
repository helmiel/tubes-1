package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type TopikArr struct {
	info [ARR_STATIC_MAX]string
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
	fmt.Scanln(&pasien.nama)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&pasien.password)

	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&pasien.umur)

	hasil = PasienFind(*arr, pasien)
	if hasil == -1 {
		PasienPush(arr, pasien)
	} else {
		fmt.Println("[info]: Pengguna sudah terdaftar")
	}
}

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
		fmt.Print("[", i+1, "] Judul:\n", f.pertanyaan.info[i].judul, "\n")

		fmt.Println("Diskusi:")
		for j := 0; j < f.pertanyaan.info[i].replies.n; j++ {
			fmt.Print("[", f.pertanyaan.info[i].replies.info[j].nama, "(", StringCapitalize(f.pertanyaan.info[i].replies.info[j].tipe), ")]: ", f.pertanyaan.info[i].replies.info[j].message, "\n")
		}
		fmt.Println("=====================")
	}
}

/* Pertanyaan F() +
 */
func PertanyaanSortAsc(p *PertanyaanArr) {
	var i, j, min_idx int
	for i = 0; i < p.n-1; i++ {
		min_idx = i
		for j = i + 1; j < p.n; j++ {
			if p.info[j].replies.n < p.info[min_idx].replies.n {
				min_idx = j
			}
		}
		var temp = p.info[i]
		p.info[i] = p.info[min_idx]
		p.info[min_idx] = temp
	}
}

func PertanyaanSortDesc(p *PertanyaanArr) {
	for i := 1; i < p.n; i++ {
		key := p.info[i]
		for j := i - 1; j > 0 && p.info[j].replies.n > key.replies.n; j-- {
			p.info[j] = p.info[j-1]
		}
		p.info[i] = key
	}
}

func Daftar() {
	PasienDaftar(&db.pasien)
}

func Login() {
	fmt.Println(`
Login sebagai
-------------
1. Pasien
2. Dokter
0. Batalkan
	`)

	var pilihan int
	for pilihan = -1; !(pilihan >= 1 && pilihan <= 2) && pilihan != 0; {
		fmt.Print("Masukan Pilihan: ")
		fmt.Scanln(&pilihan)
	}

	if pilihan != 0 {
		var nama, password string

		fmt.Print("Masukan Nama: ")
		fmt.Scanln(&nama)

		fmt.Print("Masukan Password: ")
		fmt.Scanln(&password)

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
	var loggedAsDokter bool = db.user.tipe == "DOKTER"
	var loggedAsPasien bool = db.user.tipe == "PASIEN"
	var pilihan_max = 3

	fmt.Println(`
Forum
-----
1. Lihat
2. Tambah
3. Balas`)

	if loggedAsDokter {
		pilihan_max++
		fmt.Println("4. Tools <@sp-dokter>")
	}
	fmt.Println("0. Batalkan")
	fmt.Print("\n")

	var pilihan int
	for pilihan = -1; !(pilihan >= 1 && pilihan <= pilihan_max) && pilihan != 0; {
		fmt.Print("Masukan Pilihan: ")
		fmt.Scanln(&pilihan)
	}

	if pilihan != 0 {
		if pilihan == 1 {
			ForumPrint(db.forum)
		} else if pilihan == 2 {
			if loggedAsPasien {
				var pertanyaan string
				fmt.Print("Masukan pertanyaan: ")
				ScanString(&pertanyaan)
				ForumPush(&db.forum, Pertanyaan{pasien: *db.user.pasien, judul: pertanyaan})
			} else {
				fmt.Println("[info]: Harap login sebagai pasien terlebih dahulu")
			}
		} else if pilihan == 3 {
			var reply Reply
			var i int
			if loggedAsPasien || loggedAsDokter {
				ForumPrint(db.forum)
				fmt.Print("Masukkan balasan pada forum ke: ")
				fmt.Scanln(&i)
				if i > 0 && i <= db.forum.pertanyaan.n {
					fmt.Print("Masukkan balasan anda: ")
					ScanString(&reply.message)
					reply.nama = db.user.pasien.nama
					reply.tipe = db.user.tipe
					ReplyPush(&db.forum.pertanyaan.info[i-1].replies, reply)
				} else {
					fmt.Print("[info]: Pertanyaan tidak ditemukan")
				}
			} else {
				fmt.Print("[info]: Harap login terlebih dahulu")
			}
		} else if pilihan == 4 {
			if loggedAsDokter {
				// + TOOLS
				var input int
				fmt.Println(`
Tools Dokter
------------
1. Urut / Sort pertanyaan
0. Batalkan
				`)

				for input = -1; !(input >= 1 && input <= 1) && input != 0; {
					fmt.Print("Masukkan pilihan: ")
					fmt.Scanln(&input)
				}

				if input != 0 {
					fmt.Println(`
Urut secara
-----------
1. Ascending
2. Descending
0. Batalkan
					`)

					for input = -1; !(input >= 1 && input <= 2) && input != 0; {
						fmt.Print("Masukkan pilihan: ")
						fmt.Scanln(&input)
					}

					if input == 1 {
						PertanyaanSortAsc(&db.forum.pertanyaan)
					} else if input == 2 {
						PertanyaanSortDesc(&db.forum.pertanyaan)
					}
				}
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

func ReplyPush(r *ReplyArr, x Reply) {
	if r.n < ARR_STATIC_MAX {
		r.info[r.n] = x
		r.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Balasan")
	}
}

func TopikPush(t *TopikArr, x string) {
	if t.n < ARR_STATIC_MAX {
		t.info[t.n] = x
		t.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Topik")
	}
}

func TopikSort(t *TopikArr) {
	for i := 0; i < t.n; i++ {
		for j := i + 1; j < t.n; j++ {
			if t.info[j][0] < t.info[i][0] {
				tmp := t.info[i]
				t.info[i] = t.info[j]
				t.info[j] = tmp
			}
		}
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
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "nala", message: "xxx", tipe: "PASIEN"})
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "aku", message: "xxx", tipe: "PASIEN"})
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "dia", message: "xxx", tipe: "PASIEN"})

	ForumPush(&db.forum, Pertanyaan{pasien: *db.user.pasien, judul: "apa itu kucing"})
	ReplyPush(&db.forum.pertanyaan.info[1].replies, Reply{nama: "joko", message: "xxx", tipe: "PASIEN"})

	for i := -1; i != 0; {
		Menu()
		fmt.Print("Masukkan: ")
		fmt.Scanln(&i)

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

func ScanString(buf *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*buf = scanner.Text()
}

func StringCapitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}

func Logout() {
	db.user.tipe = "LOGOUT"
}
