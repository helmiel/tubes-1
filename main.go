package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var USER_TIPE = [2]string{"PASIEN", "DOKTER"}

const ARR_STATIC_MAX int = 1024

func btoi(b bool) int {
    if b {
        return 1
    }

    return 0
}

type ArrInt struct {
    info [ARR_STATIC_MAX]int
    n int
}

type ArrString struct {
    info [ARR_STATIC_MAX]string
    n int
}

type User struct {
    tipe string
    pasien *Pasien
    dokter *Dokter
}

type Pasien struct {
    nama, password, username, NOKTP string
    umur int
}

type PasienArr struct {
    info [ARR_STATIC_MAX]Pasien
    n int
}

type Dokter struct {
    nama, username, password string
    umur int
}

type DokterArr struct {
    info [ARR_STATIC_MAX]Dokter
    n int
}

type Forum struct {
    pertanyaan PertanyaanArr
}

type Pertanyaan struct {
    judul, topik string
    pasien Pasien
    replies ReplyArr
}

type PertanyaanArr struct {
    info [ARR_STATIC_MAX]Pertanyaan
    n int
}

type Reply struct {
    username, message, tipe string
}

type ReplyArr struct {
    info [ARR_STATIC_MAX]Reply
    n int
}

/* ArrInt
*/
func ArrStringPush(a *ArrString, x string) {
    if a.n < ARR_STATIC_MAX {
        a.info[a.n] = x
        a.n++
    } else {
        fmt.Println("[info]: Gagal menambahkan String ke Array")
    }
}

func ArrStringMax(a ArrString) int {
    var idx int = 0;
    for i := 1; i < a.n; i++ {
        if (len(a.info[i]) > len(a.info[idx])) {
            idx = i;
        }
    }
    return idx
}

/* ArrInt
*/
func ArrIntPush(a *ArrInt, x int) {
    if a.n < ARR_STATIC_MAX {
        a.info[a.n] = x
        a.n++
    } else {
        fmt.Println("[info]: Gagal menambahkan Integer ke Array")
    }
}

/* Pasien
*/
func PasienFind(arr PasienArr, username, password string) int {
    var idx int = -1
    for i := 0; i < arr.n && idx == -1; i++ {
        if username == arr.info[i].username && password == arr.info[i].password {
            idx = i
        }
    }
    return idx
}

func PasienFindByUsername(arr PasienArr, username string) int {
    var idx int = -1
    for i := 0; i < arr.n && idx == -1; i++ {
        if username == arr.info[i].username {
            idx = i
        }
    }
    return idx
}

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
            found = mid
        }
    }
    return found
}

func PasienPrint(p Pasien) {
    fmt.Println(p.username, p.nama, p.umur, p.NOKTP)
}

func PasienPush(arr *PasienArr, x Pasien) {
    if arr.n < ARR_STATIC_MAX {
        arr.info[arr.n] = x
        arr.n++
    } else {
        fmt.Println("[info]: Gagal menambahkan Pasien")
    }
}

/* Selection Sort */
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

/* Dokter
*/
func DokterFind(arr DokterArr, username, password string) int {
    var idx int = -1
    for i := 0; i < arr.n && idx == -1; i++ {
        if username == arr.info[i].username && password == arr.info[i].password {
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

/* Pertanyaan
*/
func PertanyaanFind(p PertanyaanArr, topik string) ArrInt {
    var found ArrInt
    var i int
    for i = 0; i < p.n; i++ {
        if strings.ToLower(p.info[i].topik) == strings.ToLower(topik) {
            ArrIntPush(&found, i)
        }
    }
    return found
}

func PertanyaanPush(p *PertanyaanArr, x Pertanyaan) {
    if p.n < ARR_STATIC_MAX {
        p.info[p.n] = x
        p.n++
    } else {
        fmt.Println("[info]: Gagal menambahkan Pertanyaan")
    }
}

/* Selection Sort */
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

/* Insertion Sort */
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
    fmt.Println("[Judul]:", p.judul)
    fmt.Println("[Topik]:", p.topik)
    fmt.Println("[Diskusi]:")
    for j := 0; j < p.replies.n; j++ {
        fmt.Print("@", p.replies.info[j].username, " (", StringCapitalize(p.replies.info[j].tipe), "): ", p.replies.info[j].message, "\n")
    }
    fmt.Print("=====================\n\n")
}

/* Forum
*/
func ForumPrint(f Forum) {
    for i := 0; i < f.pertanyaan.n; i++ {
        fmt.Print("[", i+1, "]\n")
        PertanyaanPrint(f.pertanyaan.info[i])
    }
}

/* Relpy
*/
func ReplyPush(r *ReplyArr, x Reply) {
    if r.n < ARR_STATIC_MAX {
        r.info[r.n] = x
        r.n++
    } else {
        fmt.Println("[info]: Gagal menambahkan Balasan")
    }
}

type Database struct {
    user User
    pasien PasienArr
    dokter DokterArr
    forum Forum
}

const HEADER_BOX_GAP = 2
const HEADER_BOX_SPACE = 2
const HEADER_PILLAR_WIDTH = 3

func HeaderAlas(width int) {
    for i := 0; i < HEADER_PILLAR_WIDTH; i++ {
        fmt.Print("*")
    }
    for i := 0; i < (HEADER_BOX_SPACE / 2) + (HEADER_BOX_GAP / 2); i++ {
        fmt.Print(" ")
    }
    for i := 0; i < width; i++ {
        fmt.Print("-")
    }
    for i := 0; i < (HEADER_BOX_SPACE / 2) + (HEADER_BOX_GAP / 2); i++ {
        fmt.Print(" ")
    }
    for i := 0; i < HEADER_PILLAR_WIDTH; i++ {
        fmt.Print("*")
    }
    fmt.Println()
}

func Header(teks ArrString) {
    if (HEADER_BOX_GAP % 2 == 0 && HEADER_BOX_SPACE % 2 == 0) {
        var maks string = teks.info[ArrStringMax(teks)];
        var width int = len(maks) + HEADER_BOX_GAP + HEADER_BOX_SPACE

        // Atas
        HeaderAlas(width)

        // Tengah
        for i := 0; i < teks.n; i++ {
            var jarak_antara_teks int = (len(maks) - len(teks.info[i])) / 2
            var kurang_satu_space int /* bool */

            // cek kalo perlu +1 spasi
            if (len(maks) % 2 == 0) {
                kurang_satu_space = btoi(len(teks.info[i]) % 2 != 0)
            } else {
                kurang_satu_space = btoi(len(teks.info[i]) % 2 == 0)
            }

            // Samping
            for i := 0; i < HEADER_PILLAR_WIDTH; i++ {
                fmt.Print("*")
            }
            for j := 0; j < HEADER_BOX_SPACE + HEADER_BOX_GAP; j++ {
                fmt.Print(" ")
            }

            // Tengah
            for j := 0; j < jarak_antara_teks; j++ {
                fmt.Print(" ")
            }
            fmt.Print(teks.info[i])
            for j := 0; j < jarak_antara_teks; j++ {
                fmt.Print(" ")
            }

            // Samping
            for j := 0; j < HEADER_BOX_SPACE + HEADER_BOX_GAP + kurang_satu_space; j++ {
                fmt.Print(" ")
            }
            for i := 0; i < HEADER_PILLAR_WIDTH; i++ {
                fmt.Print("*")
            }


            fmt.Println()
        }

        // Bawah
        HeaderAlas(width)
    } else {
        fmt.Println("[info]: Gap dan space box harus genap.")
    }
}

func Menu() {
    var teks ArrString
    ArrStringPush(&teks, "Tugas Besar")
    ArrStringPush(&teks, "Algoritma Pemograman 2023")
    ArrStringPush(&teks, "Konsultasi Kesehatan")
    ArrStringPush(&teks, "Ditulis Fattan & Helmi")

    Header(teks)

    fmt.Println(`
Konsultasi Kesehatan
--------------------
1. Daftar
2. Login
3. Logout
4. Forum
0. Keluar`)
}

var db Database

func main() {
    DokterPush(&db.dokter, Dokter{username: "helmi", nama: "Helmi", password: "admin", umur: 19})
    DokterPush(&db.dokter, Dokter{username: "fattan", nama: "Fattan", password: "admin", umur: 19})

    PasienPush(&db.pasien, Pasien{username: "nala", nama: "nala", password: "nala", umur: 19, NOKTP: "12345678"})
    PasienPush(&db.pasien, Pasien{username: "sari", nama: "sari", password: "nala", umur: 19, NOKTP: "1998811"})
    PasienPush(&db.pasien, Pasien{username: "jak", nama: "jake", password: "123", umur: 19, NOKTP: "12356748"})

    db.user.tipe = USER_TIPE[0]
    db.user.pasien = &db.pasien.info[PasienFindByUsername(db.pasien, "nala")]

    PertanyaanPush(&db.forum.pertanyaan, Pertanyaan{pasien: *db.user.pasien, judul: "Apa itu Lambung", topik: "Lambung"})
    ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{username: "jak", message: "OK", tipe: "PASIEN"})
    ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{username: "nala", message: "OK?", tipe: "PASIEN"})
    ReplyPush(&db.forum.pertanyaan.info[0].replies, Reply{username: "sari", message: "OK??", tipe: "PASIEN"})

    PertanyaanPush(&db.forum.pertanyaan, Pertanyaan{pasien: *db.user.pasien, judul: "Bagaimana Cara Kucing Tidur", topik: "Kucing"})
    ReplyPush(&db.forum.pertanyaan.info[1].replies, Reply{username: "sari", message: "balon ku ada 5", tipe: "PASIEN"})

    Logout()

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
        }
    }
}

func ScanString(buf *string) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    *buf = scanner.Text()
}

func StringCapitalize(str string) string {
    return strings.Title(strings.ToLower(str))
}

/* Menu
 */

func Logout() {
    db.user.tipe = ""
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
                    if loggedAsDokter {
                        reply.username = db.user.dokter.username
                    } else {
                        reply.username = db.user.pasien.username
                    }
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
                found := PertanyaanFind(db.forum.pertanyaan, input)
                for i := 0; i < found.n; i++ {
                    fmt.Printf("[%d]\n", found.info[i] + 1)
                    PertanyaanPrint(db.forum.pertanyaan.info[found.info[i]])
                }
            } else {
                fmt.Print("[info]: Harap login terlebih dahulu")
            }
        } else if pilihan == 5 {
            if loggedAsDokter {
                var input int
                fmt.Println(`
Tools Dokter
------------
1. Urut / Sort pertanyaan
2. Cari NOKTP
0. Batalkan`)

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
4. Cari
5. Tools <@sp-dokter>
0. Batalkan`)
                    pilihan = -1
                } else if input == 1 {
                    fmt.Println(`
Urut secara
-----------
1. Ascending
2. Descending
0. Batalkan`)


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
                    var input string
                    var idx int
                    fmt.Print("Masukkan NOKTP pasien yang dicari: ")
                    fmt.Scanln(&input)
                    idx = PasienFindByNOKTPBinary(db.pasien, input)
                    if idx != -1 {
                        fmt.Print("[info]: Pasien ditemukan ")
                        PasienPrint(db.pasien.info[idx])
                    } else {
                        fmt.Print("[info]: Pasien tidak ditemukan")
                    }
                }
            }
        }
    }
}

func Daftar() {
    var pasien Pasien
    var hasil int

    fmt.Print("Masukkan Username: ")
    fmt.Scanln(&pasien.username)

    hasil = PasienFindByUsername(db.pasien, pasien.username)
    if hasil == -1 {
        fmt.Print("Masukkan Nama: ")
        fmt.Scanln(&pasien.nama)

        fmt.Print("Masukkan Umur: ")
        fmt.Scanln(&pasien.umur)

        fmt.Print("Masukkan NOKTP: ")
        fmt.Scanln(&pasien.NOKTP)

        hasil = PasienFindByNOKTPBinary(db.pasien, pasien.NOKTP)
        if hasil == -1{
            fmt.Print("Masukkan Password: ")
            fmt.Scanln(&pasien.password)

            PasienPush(&db.pasien, pasien)
            PasienSort(&db.pasien)
        } else {
            fmt.Println("[info]: Pengguna dengan NO KTP terisu sudah terdaftar")
        }
    } else {
        fmt.Println("[info]: Pengguna sudah terdaftar")
    }
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
        var username, password string

        fmt.Print("Masukan Username: ")
        fmt.Scanln(&username)

        fmt.Print("Masukan Password: ")
        fmt.Scanln(&password)

        if pilihan == 1 {
            idx := PasienFind(db.pasien, username, password)
            if idx > -1 {
                db.user.tipe = USER_TIPE[pilihan-1]
                db.user.pasien = &db.pasien.info[idx]
                fmt.Println("[info]: Selamat datang pasien", db.pasien.info[idx].nama)
            } else {
                fmt.Println("[info]: Login gagal")
            }
        } else if pilihan == 2 {
            idx := DokterFind(db.dokter, username, password)
            if idx > -1 {
                db.user.tipe = USER_TIPE[pilihan-1]
                db.user.dokter = &db.dokter.info[idx]
                fmt.Println("[info]: Selamat datang dokter", db.dokter.info[idx].nama)
            } else {
                fmt.Println("[info]: Login gagal")
            }
        }
    }
}
