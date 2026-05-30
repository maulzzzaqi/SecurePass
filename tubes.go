package main

import (
	"fmt"
	"time"
)

const akunMax = 100
const userMax = 10

type user struct {
	username string
	password string
	kumpulanAkun [akunMax]Akun
	jumlahAkun int
}

type Akun struct {
	layanan    string
	email      string
	password   string
	lastUpdate string
}

var dataAkun [MAX]Akun
var n int = 0

func inputString(pesan string) string {
	var x string
	fmt.Print(pesan)
	fmt.Scan(&x)
	return x
}

func inputTanggal() string {
	var tgl string

	tgl = time.Now().Format("02-01-2006 15:04:05")

	return tgl
}

func cekKekuatanPassword(pass string) string {
	var i, panjang, skor int
	var ch byte
	var adaHurufBesar, adaHurufKecil, adaAngka, adaSimbol bool

	panjang = len(pass)

	adaHurufBesar = false
	adaHurufKecil = false
	adaAngka = false
	adaSimbol = false

	for i = 0; i < panjang; i++ {
		ch = pass[i]

		if ch >= 'A' && ch <= 'Z' {
			adaHurufBesar = true
		} else if ch >= 'a' && ch <= 'z' {
			adaHurufKecil = true
		} else if ch >= '0' && ch <= '9' {
			adaAngka = true
		} else {
			adaSimbol = true
		}
	}

	skor = 0

	if panjang >= 8 {
		skor++
	}
	if adaHurufBesar {
		skor++
	}
	if adaHurufKecil {
		skor++
	}
	if adaAngka {
		skor++
	}
	if adaSimbol {
		skor++
	}

	if skor <= 2 {
		return "LEMAH"
	} else if skor <= 4 {
		return "SEDANG"
	}
	return "KUAT"
}

func tambahUser() {
	var i int
	var newUsername string
	var isDuplicate bool

	isDuplicate = false

	if totalUser < userMax {
		fmt.Println("\n=== TAMBAH USER ===")

		newUsername = inputString("Username      :")

		for i = 0; i < totalUser; i++ {
			if dataUser[i].username == newUsername {
				isDuplicate = true
			}
		}

		if isDuplicate {
			fmt.Println("Gagal! Username sudah digunakan. Silakan coba yang lain.")
		} else {
			dataUser[totalUser].username = newUsername
			dataUser[totalUser].password = inputString("Password      :")

			totalUser++

			fmt.Println("User berhasil ditambahkan!")
		}
	} else {
		fmt.Println("Data penuh! Tidak bisa menambah user.")
	}
}

func loginUser() {
	var inputUsername, inputPassword string
	var i int
	var isFound bool

	isFound = false

	if totalUser == 0 {
		fmt.Println("\n=== LOGIN USER ===")
		fmt.Println("Error! Tidak ada User yang terdaftar!")
		fmt.Println("Silahkan lakukan Register terlebih dahulu!")
		tambahUser()
	} else {
		fmt.Println("\n=== LOGIN USER ===")
		inputUsername = inputString("Username    :")
		inputPassword = inputString("Password    :")

		for i = 0; i < totalUser; i++ {
			if dataUser[i].username == inputUsername && dataUser[i].password == inputPassword {
				currentUserIndex = i
				isFound = true
			}
		}

		if isFound {
			fmt.Println("Login berhasil!")
		} else {
			fmt.Println("Login gagal! Username atau password salah.")
		}
	}
}

func tambahAkun() {
	if n < MAX {
		fmt.Println("\n=== TAMBAH AKUN ===")
		dataAkun[n].layanan = inputString("Nama Layanan      : ")
		dataAkun[n].email = inputString("Email/Username    : ")
		dataAkun[n].password = inputString("Password          : ")
		dataAkun[n].lastUpdate = inputTanggal()

		n++
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Data penuh! Tidak bisa menambah akun.")
	}
}

func tampilkanAkun() {
	var i int
	fmt.Println("\n=== DAFTAR AKUN ===")

	if n == 0 {
		fmt.Println("Belum ada data akun.")
	} else {
		for i = 0; i < n; i++ {
			fmt.Println("----------------------------")
			fmt.Println("No          :", i+1)
			fmt.Println("Layanan     :", dataAkun[i].layanan)
			fmt.Println("Email       :", dataAkun[i].email)
			fmt.Println("Password    :", dataAkun[i].password)
			fmt.Println("Last Update :", dataAkun[i].lastUpdate)
			fmt.Println("Kekuatan    :", cekKekuatanPassword(dataAkun[i].password))
		}
		fmt.Println("----------------------------")
	}
}

func ubahAkun() {
	var idx, i int
	fmt.Println("\n=== UBAH AKUN ===")

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa diubah.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin diubah: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			i = idx - 1

			fmt.Println("\nData lama:")
			fmt.Println("Layanan :", dataAkun[i].layanan)
			fmt.Println("Email   :", dataAkun[i].email)
			fmt.Println("Pass    :", dataAkun[i].password)
			fmt.Println("Update  :", dataAkun[i].lastUpdate)

			dataAkun[i].layanan = inputString("Layanan baru   : ")
			dataAkun[i].email = inputString("Email baru     : ")
			dataAkun[i].password = inputString("Password baru  : ")
			dataAkun[i].lastUpdate = inputTanggal()

			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

func hapusAkun() {
	var idx, i, pos int
	fmt.Println("\n=== HAPUS AKUN ===")

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa dihapus.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin dihapus: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			pos = idx - 1

			for i = pos; i < n-1; i++ {
				dataAkun[i] = dataAkun[i+1]
			}

			n--
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

// searching 
// Sequential Search
func sequentialSearch(layanan string) int {
	var pos, i int
	pos = -1

	for i = 0; i < n; i++ {
		if dataAkun[i].layanan == layanan {
			pos = i
		}
	}
	return pos
}

// Binary Search (data harus sudah diurutkan alfabet)
func binarySearch(layanan string) int {
	var kiri, kanan, pos, tengah int

	kiri = 0
	kanan = n - 1
	pos = -1

	for kiri <= kanan && pos == -1 {
		tengah = (kiri + kanan) / 2

		if dataAkun[tengah].layanan == layanan {
			pos = tengah
		} else if dataAkun[tengah].layanan < layanan {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return pos
}

func menuCari() {
	var pilih, pos int
	var layanan string

	fmt.Println("\n=== MENU PENCARIAN ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (harus sudah sorting alfabet)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	layanan = inputString("Masukkan nama layanan yang dicari: ")

	if pilih == 1 {
		pos = sequentialSearch(layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", dataAkun[pos].layanan)
			fmt.Println("Email   :", dataAkun[pos].email)
			fmt.Println("Pass    :", dataAkun[pos].password)
			fmt.Println("Update  :", dataAkun[pos].lastUpdate)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	} else if pilih == 2 {
		pos = binarySearch(layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", dataAkun[pos].layanan)
			fmt.Println("Email   :", dataAkun[pos].email)
			fmt.Println("Pass    :", dataAkun[pos].password)
			fmt.Println("Update  :", dataAkun[pos].lastUpdate)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

//SORTING 
// Selection Sort berdasarkan nama layanan alfabet
func selectionSortNama() {
	var i, j, min int
	var temp Akun

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if dataAkun[j].layanan < dataAkun[min].layanan {
				min = j
			}
		}

		temp = dataAkun[i]
		dataAkun[i] = dataAkun[min]
		dataAkun[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama layanan (A-Z).")
}

// Insertion Sort berdasarkan tanggal update
func insertionSortTanggal() {
	var i, j int
	var key Akun

	for i = 1; i < n; i++ {
		key = dataAkun[i]
		j = i - 1

		for j >= 0 && dataAkun[j].lastUpdate > key.lastUpdate {
			dataAkun[j+1] = dataAkun[j]
			j--
		}

		dataAkun[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tanggal update.")
}

func menuSort() {
	var pilih int

	fmt.Println("\n=== MENU SORTING ===")
	fmt.Println("1. Selection Sort (Nama layanan A-Z)")
	fmt.Println("2. Insertion Sort (Tanggal update)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortNama()
	} else if pilih == 2 {
		insertionSortTanggal()
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

func statistik() {
	var lemah, kuat, sedang, i int
	var k string

	fmt.Println("\n=== STATISTIK SECUREPASS ===")
	fmt.Println("Total akun tersimpan:", n)

	lemah = 0
	sedang = 0
	kuat = 0

	for i = 0; i < n; i++ {
		k = cekKekuatanPassword(dataAkun[i].password)

		if k == "LEMAH" {
			lemah++
		} else if k == "SEDANG" {
			sedang++
		} else {
			kuat++
		}
	}

	fmt.Println("Password Lemah :", lemah)
	fmt.Println("Password Sedang:", sedang)
	fmt.Println("Password Kuat  :", kuat)
}

func menuUtama() {
	fmt.Println("\n===== SECUREPASS MENU =====")
	fmt.Println("1. Tambah Akun")
	fmt.Println("2. Tampilkan Semua Akun")
	fmt.Println("3. Ubah Akun")
	fmt.Println("4. Hapus Akun")
	fmt.Println("5. Cari Akun")
	fmt.Println("6. Urutkan Data")
	fmt.Println("7. Statistik")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var pilih int
	pilih = -1

	for pilih != 0 {
		menuUtama()
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahAkun()
		} else if pilih == 2 {
			tampilkanAkun()
		} else if pilih == 3 {
			ubahAkun()
		} else if pilih == 4 {
			hapusAkun()
		} else if pilih == 5 {
			menuCari()
		} else if pilih == 6 {
			menuSort()
		} else if pilih == 7 {
			statistik()
		} else if pilih == 0 {
			fmt.Println("Terima kasih sudah menggunakan SecurePass.")
		} else {
			fmt.Println("Menu tidak valid!")
		}
	}
}