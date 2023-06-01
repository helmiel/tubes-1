package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ARR_STATIC_MAX int = 1024

type ArrInt struct {
	info [ARR_STATIC_MAX]int
	n    int
}

type Pasien struct {
	nama, password string
	umur           int
	NOKTP          string
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
	judul, topik string
	pasien       Pasien
	replies      ReplyArr
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

	fmt.Print("Masukkan NOKTP: ")
	fmt.Scanln(&pasien.NOKTP)

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

func ForumPrint(f Forum) {
	for i := 0; i < f.pertanyaan.n; i++ {
		fmt.Print("[", i+1, "]")
		PertanyaanPrint(f.pertanyaan.info[i])
	}
}

/* Pertanyaan F() +
 */

func PertanyaanPush(p *PertanyaanArr, x Pertanyaan) {
	if p.n < ARR_STATIC_MAX {
		p.info[p.n] = x
		p.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Pertanyaan")
	}
}

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
		j := i - 1
		for ; j > 0 && p.info[j].replies.n > key.replies.n; j-- {
			p.info[j] = p.info[j-1]
		}
		p.info[j+1] = key
	}
}

func PertanyaanPrint(p Pertanyaan) {
	fmt.Print("Judul: ", p.judul, "\n")
	fmt.Println("[Topik]: ", p.topik)

	fmt.Println("Diskusi:")
	for j := 0; j < p.replies.n; j++ {
		fmt.Print("[", p.replies.info[j].nama, "(", StringCapitalize(p.replies.info[j].tipe), ")]: ", p.replies.info[j].message, "\n")
	}
	fmt.Println("=====================")
}

func ArrIntPush(a *ArrInt, x int) {
	if a.n < ARR_STATIC_MAX {
		a.info[a.n] = x
		a.n++
	} else {
		fmt.Println("[info]: Gagal menambahkan Integer ke Array")
	}
}

// Untuk asien ~ Sequential Search
func PertanyaanFindByTopikSequential(p PertanyaanArr, topik string) ArrInt {
	var found ArrInt
	var i int
	for i = 0; i < p.n; i++ {
		if p.info[i].topik == topik {
			ArrIntPush(&found, i)
		}
	}
	return found
}

// Untuk dokter ~ Binary Search
func PasienFindByNOKTPBinary(p PasienArr, NOKTP string) int {
	var kr, kn, mid int
	var found = -1
	kr = 0
	kn = p.n - 1
	for kr <= kn && found == -1 {
		mid = (kr + kn) / 2
		if NOKTP > p.info[mid].NOKTP {
			kr = mid + 1
		} else if NOKTP < p.info[mid].NOKTP {
			kn = mid - 1
		} else {
			// ArrIntPush(&found, p.info[mid])
			found = mid
		}
	}
	return found
}

func PasienSort(p *PasienArr) {
	var i, j, min_idx int
	for i = 0; i < p.n-1; i++ {
		min_idx = i
		for j = i + 1; j < p.n; j++ {
			if p.info[j].NOKTP < p.info[min_idx].NOKTP {
				min_idx = j
			}
		}
		if min_idx != i {
			var temp = p.info[i]
			p.info[i] = p.info[min_idx]
			p.info[min_idx] = temp
		}
	}
}

func PasienPrint(p Pasien) {
	fmt.Println(p.nama, p.umur, p.NOKTP)
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
	var pilihan_max = 4

	fmt.Println(`
Forum
-----
1. Lihat
2. Tambah
3. Balas
4. Cari`)

	if loggedAsDokter {
		pilihan_max++
		fmt.Println("5. Tools <@sp-dokter>")
	}
	fmt.Println("0. Batalkan")
	fmt.Print("\n")

	var pilihan int
	for pilihan = -1; !(pilihan >= 1 && pilihan <= pilihan_max) && pilihan != 0; {
		fmt.Print("Masukan Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			ForumPrint(db.forum)
		} else if pilihan == 2 {
			if loggedAsPasien {
				var pertanyaan, topik string
				fmt.Print("Masukan pertanyaan: ")
				ScanString(&pertanyaan)
				fmt.Print("Masukkan topik: ")
				fmt.Scanln(&topik)
				PertanyaanPush(&db.forum.pertanyaan, Pertanyaan{pasien: *db.user.pasien, judul: pertanyaan, topik: topik})
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
			if loggedAsPasien || loggedAsDokter {
				var input string
				fmt.Print("Masukkan topik yang akan dicari: ")
				fmt.Scanln(&input)
				found := PertanyaanFindByTopikSequential(db.forum.pertanyaan, input)
				for i := 0; i < found.n; i++ {
					PertanyaanPrint(db.forum.pertanyaan.info[found.info[i]])
				}
			} else {
				fmt.Print("[info]: Harap login terlebih dahulu")
			}
		} else if pilihan == 5 {
			if loggedAsDokter {
				// + TOOLS
				var input int
				fmt.Println(`
Tools Dokter
------------
1. Urut / Sort pertanyaan
2. Cari NOKTP
0. Batalkan
       `)

				for input = -1; !(input >= 1 && input <= 2) && input != 0; {
					fmt.Print("Masukkan pilihan: ")
					fmt.Scanln(&input)
				}

				if input == 0 {
					fmt.Println(`
Forum
-----
1. Lihat
2. Tambah
3. Balas
4. Cari`)
					pilihan = -1
				} else if input == 1 {
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
				} else if input == 2 {
					PasienSort(&db.pasien)
					var input string
					var idx int
					fmt.Print("Masukkan NOKTP pasien yang dicari: ")
					fmt.Scanln(&input)
					idx = PasienFindByNOKTPBinary(db.pasien, input)
					if idx != -1 {
						fmt.Print("[Pasien ditemukan]: ")
						PasienPrint(db.pasien.info[idx])
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

var db Database

func main() {
	DokterPush(&db.dokter, Dokter{nama: "Helmi", password: "admin", umur: 19})
	DokterPush(&db.dokter, Dokter{nama: "Fattan", password: "admin", umur: 19})

	PasienPush(&db.pasien, Pasien{nama: "nala", password: "nala", umur: 19, NOKTP: "12345678"})
	PasienPush(&db.pasien, Pasien{nama: "aku", password: "123", umur: 19, NOKTP: "12356748"})

	db.user.tipe = USER_TIPE[0]
	db.user.pasien = &db.pasien.info[PasienFind(db.pasien, Pasien{nama: "nala", password: "nala", umur: 19})]

	PertanyaanPush(&db.forum.pertanyaan, Pertanyaan{pasien: *db.user.pasien, judul: "apa itu lambung", topik: "Lambung"})
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "nala", message: "xxx", tipe: "PASIEN"})
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "aku", message: "xxx", tipe: "PASIEN"})
	ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{nama: "dia", message: "xxx", tipe: "PASIEN"})

	PertanyaanPush(&db.forum.pertanyaan, Pertanyaan{pasien: *db.user.pasien, judul: "apa itu kucing", topik: "Lambung"})
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
