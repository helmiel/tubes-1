package main

import (
    "errors"
    "fmt"
)

const ARRAY_MAX__STATIC int = 1024

type Array[T any] struct {
    val [ARRAY_MAX__STATIC]T
    len int
}

func (a *Array[T]) Init() {
    a.len = 0
}

func (a Array[T]) Get(idx int) (T, error) {
    var ret T
    if idx >= a.len || idx < 0 {
        return ret, errors.New("Array out of bounds")
    }

    ret = a.val[idx]
    return ret, nil
}

func (a *Array[T]) Push(v T) error {
    if a.len >= ARRAY_MAX__STATIC {
        return errors.New("Array cannot hold any more items")
    }

    a.val[a.len] = v
    a.len++
    return nil
}

func (a *Array[T]) Pop() (T, error) {
    var ret T
    if a.len <= 0 {
        return ret, errors.New("Array cannot delete any items")
    }

    ret = a.val[a.len-1]
    a.len--
    return ret, nil
}

func (a Array[T]) Find(f func(v T, i int) bool) int {
    var idx = -1

    for i := 0; i < a.len; i++ {
        if f(a.val[i], i) {
            idx = i
        }
    }

    return idx
}

func (a Array[T]) Each(f func(v T, i int)) {
    for i := 0; i < a.len; i++ {
        f(a.val[i], i)
    }
}

func main() {
    var orang Orang
    var db Database

    for i := -1; i != 0; {
        Menu()
        fmt.Print("Masukan: ")
        fmt.Scan(&i)

        if i == 1 {
            fmt.Println(`
Daftarkan dirimu sebagai
1) Pasien
2) Dokter
            `)

            var idx int

            fmt.Print("Masukan Pilihan: ")
            fmt.Scan(&idx)

            // disarankan tidak pake break -pras
            if idx >= 1 && idx <= len(ORANG_TIPE) {
                var nama string
                var password string

                fmt.Print("Masukan Nama: ")
                fmt.Scan(&nama)

                fmt.Print("Masukan Password: ")
                fmt.Scan(&password)

                orang.Init(ORANG_TIPE[idx-1], nama, password)
                db.orang.Push(orang)
                fmt.Println("Berhasil terdaftar")
            }
        } else if i == 2 {
            fmt.Println(`
Login sebagai
1) Pasien
2) Dokter
            `)

            var pilihan int

            fmt.Print("Masukan Pilihan: ")
            fmt.Scan(&pilihan)

            if pilihan >= 1 && pilihan <= len(ORANG_TIPE) {
                var nama, password string

                fmt.Print("Masukan Nama: ")
                fmt.Scan(&nama)

                fmt.Print("Masukan Password: ")
                fmt.Scan(&password)

                idx := db.orang.Find(func(v Orang, i int) bool {
                    return v.tipe == ORANG_TIPE[pilihan-1] && v.nama == nama && v.password == password
                })

                if idx == -1 {
                    fmt.Println("Gagal login")
                } else {
                    fmt.Println("Berhasil login")
                    orang = db.orang.val[idx]
                }
            }
        }
    }
}

var ORANG_TIPE = [2]string{"PASIEN", "DOKTER"}

type Orang struct {
    tipe, nama, password string
}

func (o *Orang) Init(tipe, nama, password string) {
    o.nama = nama
    o.tipe = tipe
    o.password = password
}

type Pertanyaan struct {
    judul   string
    topik   Array[string]
    balasan Array[Balasan]
}

func (p *Pertanyaan) Init(judul string, topik Array[string]) {
    p.judul = judul
    p.topik = topik
}

type Balasan struct {
    orang Orang
    isi   string
}

func (b *Balasan) Init(o Orang, isi string) {
    b.orang = o
    b.isi = isi
}

type Database struct {
    orang Array[Orang]
    forum Array[Pertanyaan]
}

func Menu() {
    fmt.Println(`
Konsultasi Kesehatan
--------------------
1) Daftar
2) Login
3) Forum
0) Keluar
    `)
}
