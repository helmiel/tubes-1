package main

import (
	"fmt"
)

type pengguna struct {
	password     string
	username     string
	umur         int
	jenisKelamin string
	noTelepon    int
	nama         string
	status       string
	teman        []string
}

type arrpengguna struct {
	arr [NMAX]pengguna
	n   int
}

const NMAX = 10

var user pengguna

func main() {
	var arr_pengguna arrpengguna
	var berhasil_login, berhasil_registrasi bool

	var pilih int
	fmt.Println("Sosial media")
	fmt.Println("1) Registrasi")
	fmt.Println("2) Login")
	fmt.Println("3) Exit")
	fmt.Println("Masukkan pilihan Anda")
	fmt.Scan(&pilih)

	for pilih != 3 {
		if pilih == 1 {
			registrasi(&arr_pengguna, &berhasil_registrasi)
			if berhasil_registrasi {
				fmt.Println("Anda berhasil registrasi")
			} else {
				fmt.Println("Anda gagal melakukan registrasi")
			}
		} else if pilih == 2 {
			berhasil_login = login(arr_pengguna)
			if berhasil_login {
				fmt.Println("Anda berhasil login")
				menuUtama(&arr_pengguna)
			} else {
				fmt.Println("Username atau password tidak cocok")
			}
		}
		fmt.Println("Sosial media")
		fmt.Println("1) Registrasi")
		fmt.Println("2) Login")
		fmt.Println("3) Exit")
		fmt.Println("Masukkan pilihan Anda")
		fmt.Scan(&pilih)
	}

}

func login(array arrpengguna) bool {
	var akun pengguna
	var berhasil bool = false

	fmt.Println("Silakan Login")
	fmt.Println("Masukkan Username")
	fmt.Scan(&akun.username)

	fmt.Println("Masukkan Password")
	fmt.Scan(&akun.password)

	for i := 0; i < array.n; i++ {
		if akun.username == array.arr[i].username && akun.password == array.arr[i].password {
            user = array.arr[i]
			berhasil = true
		}
	}
	return berhasil
}

func registrasi(array *arrpengguna, berhasil *bool) {
	var pengguna_baru pengguna
	fmt.Println("Silakan Registrasi")

	fmt.Println("Masukkan Umur")
	fmt.Scan(&pengguna_baru.umur)
	umur := pengguna_baru.umur

	if umur >= 17 {
		*berhasil = true

		fmt.Println("Nama:")
		fmt.Scan(&pengguna_baru.nama)

		fmt.Println("Jenis Kelamin:")
		fmt.Scan(&pengguna_baru.jenisKelamin)

		fmt.Println("No Telepon:")
		fmt.Scan(&pengguna_baru.noTelepon)

		fmt.Println("Masukkan Username")
		fmt.Scan(&pengguna_baru.username)

		fmt.Println("Masukkan Password")
		fmt.Scan(&pengguna_baru.password)

		array.arr[array.n] = pengguna_baru
        user = array.arr[array.n]
		array.n += 1
	} else {
		*berhasil = false

		fmt.Println("Maaf Anda Belum Cukup Umur")
	}
}

func menuUtama(arr_pengguna *arrpengguna) {

	var pilihan int
    fmt.Println("Menu Utama")
    fmt.Println("1) Lihat Status")
    fmt.Println("2) Lihat Status")
    fmt.Println("3) Komentar Status")
    fmt.Println("2) Lihat Status")
    fmt.Println("4) Tambah Teman")
    fmt.Println("5) Hapus Teman")
    fmt.Println("6) Edit Profil")
    fmt.Println("7) Lihat Data Teman")
    fmt.Println("8) Cari Pengguna")
    fmt.Println("9) Keluar")
    fmt.Println("Masukkan pilihan Anda")
    fmt.Scan(&pilihan)

	for pilihan != 9 {
		if pilihan == 1 {
			lihatStatus(arr_pengguna)
        } else if pilihan == 1 {
            tambahStatus(arr_pengguna)
		} else if pilihan == 3 {
			komentarStatus(arr_pengguna)
		} else if pilihan == 4 {
			tambahTeman(arr_pengguna)
		} else if pilihan == 5 {
			hapusTeman(arr_pengguna)
		} else if pilihan == 6 {
			editProfil(arr_pengguna)
		} else if pilihan == 7 {
			lihatDataTeman(arr_pengguna)
		} else if pilihan == 7 {
			cariPengguna(arr_pengguna)
		} else if pilihan == 8 {
			cariPengguna(arr_pengguna)
		} else {
			fmt.Println("Pilihan tidak valid")
		}

		fmt.Println("Menu Utama")
		fmt.Println("1) Lihat Status")
		fmt.Println("2) Tambah Status")
		fmt.Println("3) Komentar Status")
		fmt.Println("2) Lihat Status")
		fmt.Println("4) Tambah Teman")
		fmt.Println("5) Hapus Teman")
		fmt.Println("6) Edit Profil")
		fmt.Println("7) Lihat Data Teman")
		fmt.Println("8) Cari Pengguna")
		fmt.Println("9) Keluar")
		fmt.Println("Masukkan pilihan Anda")
		fmt.Scan(&pilihan)

	}
}

func lihatStatus(arr_pengguna *arrpengguna) {
	for i := 0; i < arr_pengguna.n; i++ {
		fmt.Println("Username:", arr_pengguna.arr[i].username)
		fmt.Println("Status:", arr_pengguna.arr[i].status)
		fmt.Println("----------------------------")
	}
}

func tambahStatus(arr_pengguna *arrpengguna) {
	idx := find(*arr_pengguna, user.username)
	if idx == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	var status string
	fmt.Println("Masukkan Status Anda")
	fmt.Scan(&status)

    arr_pengguna.arr[idx].status = status
}

func komentarStatus(arr_pengguna *arrpengguna) {
    var username string
    fmt.Println("Masukkan nama user yang ingin dikomentari")
    fmt.Scan(&username)

	idx := find(*arr_pengguna, username)
	if idx == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	var komentar string
	fmt.Println("Masukkan komentar Anda")
	fmt.Scan(&komentar)

    arr_pengguna.arr[idx].status += "\nKomentar (" + user.username + "): " + komentar
}

func tambahTeman(arr_pengguna *arrpengguna) {
	idxPengguna := find(*arr_pengguna, user.username)
	if idxPengguna == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	var usernameTeman string
	fmt.Println("Masukkan username teman yang ingin ditambahkan")
	fmt.Scan(&usernameTeman)

	idxTeman := find(*arr_pengguna, usernameTeman)
	if idxTeman == -1 {
		fmt.Println("Teman tidak ditemukan")
		return
	}

	// Check if the friend is already in the friend list
	for _, friend := range arr_pengguna.arr[idxPengguna].teman {
		if friend == usernameTeman {
			fmt.Println("Teman sudah ada dalam daftar teman")
			return
		}
	}

	arr_pengguna.arr[idxPengguna].teman = append(arr_pengguna.arr[idxPengguna].teman, usernameTeman)
	fmt.Println("Teman berhasil ditambahkan")
}

func hapusTeman(arr_pengguna *arrpengguna) {
	idxPengguna := find(*arr_pengguna, user.username)
	if idxPengguna == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	if len(arr_pengguna.arr[idxPengguna].teman) == 0 {
		fmt.Println("Anda tidak memiliki teman")
		return
	}

	var usernameTeman string
	fmt.Println("Masukkan username teman yang ingin dihapus")
	fmt.Scan(&usernameTeman)

	idxTeman := -1
	for i := 0; i < len(arr_pengguna.arr[idxPengguna].teman); i++ {
		if arr_pengguna.arr[idxPengguna].teman[i] == usernameTeman {
			idxTeman = i
		}
	}

	if idxTeman == -1 {
		fmt.Println("Teman tidak ditemukan")
		return
	}

	// Shift the elements to the left to remove the friend
	for i := idxTeman; i < len(arr_pengguna.arr[idxPengguna].teman)-1; i++ {
		arr_pengguna.arr[idxPengguna].teman[i] = arr_pengguna.arr[idxPengguna].teman[i+1]
	}

	// Reduce the length of the friend list by 1
	arr_pengguna.arr[idxPengguna].teman = arr_pengguna.arr[idxPengguna].teman[:len(arr_pengguna.arr[idxPengguna].teman)-1]
	fmt.Println("Teman berhasil dihapus")
}

func editProfil(arr_pengguna *arrpengguna) {
	idx := find(*arr_pengguna, user.username)
	if idx == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	fmt.Println("Masukkan data baru:")
	fmt.Println("Nama:")
	fmt.Scan(&arr_pengguna.arr[idx].nama)
	fmt.Println("Jenis Kelamin:")
	fmt.Scan(&arr_pengguna.arr[idx].jenisKelamin)
	fmt.Println("No Telepon:")
	fmt.Scan(&arr_pengguna.arr[idx].noTelepon)

	fmt.Println("Profil berhasil diubah")
}

func lihatDataTeman(arr_pengguna *arrpengguna) {
	idx := find(*arr_pengguna, user.username)
	if idx == -1 {
		fmt.Println("Pengguna tidak ditemukan")
		return
	}

	if len(arr_pengguna.arr[idx].teman) == 0 {
		fmt.Println("Anda tidak memiliki teman")
		return
	}

	// Sort the friends by username
	sortedFriends := sortTeman(arr_pengguna.arr[idx].teman)

	fmt.Println("Data teman Anda:")
	for i := 0; i < len(sortedFriends); i++ {
		friend := sortedFriends[i]
		friendIdx := find(*arr_pengguna, friend)
		if friendIdx != -1 {
			fmt.Println("Username:", arr_pengguna.arr[friendIdx].username)
			fmt.Println("Nama:", arr_pengguna.arr[friendIdx].nama)
			fmt.Println("Jenis Kelamin:", arr_pengguna.arr[friendIdx].jenisKelamin)
			fmt.Println("No Telepon:", arr_pengguna.arr[friendIdx].noTelepon)
			fmt.Println("Umur:", arr_pengguna.arr[friendIdx].umur)
			fmt.Println("----------------------------")
		}
	}
}
func sortTeman(friends []string) []string {
	sortedFriends := make([]string, len(friends))
	copy(sortedFriends, friends)

	for i := 1; i < len(sortedFriends); i++ {
		key := sortedFriends[i]
		j := i - 1
		for j >= 0 && sortedFriends[j] > key {
			sortedFriends[j+1] = sortedFriends[j]
			j--
		}
		sortedFriends[j+1] = key
	}

	return sortedFriends
}

func cariPengguna(arr_pengguna *arrpengguna) {
	var keyword string
	fmt.Println("Masukkan Username pencarian")
	fmt.Scan(&keyword)

	var found bool
	for i := 0; i < arr_pengguna.n; i++ {
		if arr_pengguna.arr[i].username == keyword || arr_pengguna.arr[i].nama == keyword {
			fmt.Println("Username:", arr_pengguna.arr[i].username)
			fmt.Println("Nama:", arr_pengguna.arr[i].nama)
			fmt.Println("Jenis Kelamin:", arr_pengguna.arr[i].jenisKelamin)
			fmt.Println("No Telepon:", arr_pengguna.arr[i].noTelepon)
			fmt.Println("Umur:", arr_pengguna.arr[i].umur)
			fmt.Println("----------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Pengguna tidak ditemukan")
	}
}

func find(arr arrpengguna, username string) int {
	var idx int = -1
	for i := 0; i < arr.n; i++ {
		if arr.arr[i].username == username {
			idx = i
			break
		}
	}
	return idx
}
