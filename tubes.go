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

type tabUsers [userMax]user

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
	var jumlahAkun int = dataUser[currentUserIndex].jumlahAkun

	if jumlahAkun < akunMax {
		fmt.Println("\n=== TAMBAH AKUN ===")

		dataUser[currentUserIndex].kumpulanAkun[jumlahAkun].layanan = inputString("Nama Layanan      : ")
		dataUser[currentUserIndex].kumpulanAkun[jumlahAkun].email = inputString("Email/Username    : ")
		dataUser[currentUserIndex].kumpulanAkun[jumlahAkun].password = inputString("Password          : ")
		dataUser[currentUserIndex].kumpulanAkun[jumlahAkun].lastUpdate = inputTanggal()

		dataUser[currentUserIndex].jumlahAkun++
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Data penuh! Tidak bisa menambah akun.")
	}
}

func tampilkanAkun() {
	var i int
	var n int = dataUser[currentUserIndex].jumlahAkun
	fmt.Println("\n=== DAFTAR AKUN ===")

	if n == 0 {
		fmt.Println("Belum ada data akun.")
	} else {
		for i = 0; i < n; i++ {
			fmt.Println("----------------------------")
			fmt.Println("No          :", i+1)
			fmt.Println("Layanan     :", dataUser[currentUserIndex].kumpulanAkun[i].layanan)
			fmt.Println("Email       :", dataUser[currentUserIndex].kumpulanAkun[i].email)
			fmt.Println("Last Update :", dataUser[currentUserIndex].kumpulanAkun[i].lastUpdate)
		}
		fmt.Println("----------------------------")
	}
}

func ubahAkun() {
	var idx, i int
	var n int = dataUser[currentUserIndex].jumlahAkun
	fmt.Println("\n=== UBAH AKUN ===")
	tampilkanAkun()

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa diubah.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin diubah: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			i = idx - 1

			fmt.Println("\nData lama:")
			fmt.Println("Layanan :", dataUser[currentUserIndex].kumpulanAkun[i].layanan)
			fmt.Println("Email   :", dataUser[currentUserIndex].kumpulanAkun[i].email)
			fmt.Println("Pass    :", dataUser[currentUserIndex].kumpulanAkun[i].password)
			fmt.Println("Update  :", dataUser[currentUserIndex].kumpulanAkun[i].lastUpdate)

			dataUser[currentUserIndex].kumpulanAkun[i].layanan = inputString("Layanan baru   : ")
			dataUser[currentUserIndex].kumpulanAkun[i].email = inputString("Email baru     : ")
			dataUser[currentUserIndex].kumpulanAkun[i].password = inputString("Password baru  : ")
			dataUser[currentUserIndex].kumpulanAkun[i].lastUpdate = inputTanggal()

			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

func hapusAkun() {
	var idx, i, pos int
	var n int = dataUser[currentUserIndex].jumlahAkun
	fmt.Println("\n=== HAPUS AKUN ===")
	tampilkanAkun()

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa dihapus.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin dihapus: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			pos = idx - 1

			for i = pos; i < n-1; i++ {
				dataUser[currentUserIndex].kumpulanAkun[i] = dataUser[currentUserIndex].kumpulanAkun[i+1]
			}

			dataUser[currentUserIndex].jumlahAkun--
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

// searching 
// Sequential Search (layanan)
func sequentialSearch(layanan string) int {
	var pos, i int
	pos = -1
	var n int = dataUser[currentUserIndex].jumlahAkun

	for i = 0; i < n; i++ {
		if dataUser[currentUserIndex].kumpulanAkun[i].layanan == layanan {
			pos = i
		}
	}
	return pos
}

// Binary Search (data harus sudah diurutkan alfabet) (username atau email)
func binarySearch(email string) int {
	var kiri, kanan, pos, tengah int
	var n int = dataUser[currentUserIndex].jumlahAkun

	kiri = 0
	kanan = n - 1
	pos = -1

	for kiri <= kanan && pos == -1 {
		tengah = (kiri + kanan) / 2

		if dataUser[currentUserIndex].kumpulanAkun[tengah].email == email {
			pos = tengah
		} else if dataUser[currentUserIndex].kumpulanAkun[tengah].email < email {
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
	fmt.Println("1. Cari berdasarkan layanan")
	fmt.Println("2. Cari berdasarkan email")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	layanan = inputString("Masukkan username/email yang dicari: ")

	if pilih == 1 {
		pos = sequentialSearch(layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", dataUser[currentUserIndex].kumpulanAkun[pos].layanan)
			fmt.Println("Email   :", dataUser[currentUserIndex].kumpulanAkun[pos].email)
			fmt.Println("Pass    :", dataUser[currentUserIndex].kumpulanAkun[pos].password)
			fmt.Println("Update  :", dataUser[currentUserIndex].kumpulanAkun[pos].lastUpdate)
			fmt.Println("Kekuatan:", cekKekuatanPassword(dataUser[currentUserIndex].kumpulanAkun[pos].password))
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	} else if pilih == 2 {
		pos = binarySearch(layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", dataUser[currentUserIndex].kumpulanAkun[pos].layanan)
			fmt.Println("Email   :", dataUser[currentUserIndex].kumpulanAkun[pos].email)
			fmt.Println("Pass    :", dataUser[currentUserIndex].kumpulanAkun[pos].password)
			fmt.Println("Update  :", dataUser[currentUserIndex].kumpulanAkun[pos].lastUpdate)
			fmt.Println("Kekuatan:", cekKekuatanPassword(dataUser[currentUserIndex].kumpulanAkun[pos].password))
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
	var n int = dataUser[currentUserIndex].jumlahAkun
	var temp Akun

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if dataUser[currentUserIndex].kumpulanAkun[j].layanan < dataUser[currentUserIndex].kumpulanAkun[min].layanan {
				min = j
			}
		}

		temp = dataUser[currentUserIndex].kumpulanAkun[i]
		dataUser[currentUserIndex].kumpulanAkun[i] = dataUser[currentUserIndex].kumpulanAkun[min]
		dataUser[currentUserIndex].kumpulanAkun[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama layanan (A-Z).")
}

// Insertion Sort berdasarkan tanggal update
func insertionSortTanggal() {
	var i, j int
	var n int = dataUser[currentUserIndex].jumlahAkun
	var key Akun

	for i = 1; i < n; i++ {
		key = dataUser[currentUserIndex].kumpulanAkun[i]
		j = i - 1

		for j >= 0 && dataUser[currentUserIndex].kumpulanAkun[j].lastUpdate > key.lastUpdate {
			dataUser[currentUserIndex].kumpulanAkun[j+1] = dataUser[currentUserIndex].kumpulanAkun[j]
			j--
		}

		dataUser[currentUserIndex].kumpulanAkun[j+1] = key
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
	var n int = dataUser[currentUserIndex].jumlahAkun

	fmt.Println("\n=== STATISTIK SECUREPASS ===")
	fmt.Println("Total akun tersimpan:", n)

	lemah = 0
	sedang = 0
	kuat = 0

	for i = 0; i < n; i++ {
		k = cekKekuatanPassword(dataUser[currentUserIndex].kumpulanAkun[i].password)

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
	fmt.Printf("===== Selamat Datang, %s! ====\n", dataUser[currentUserIndex].username)
	fmt.Println("1. Tambah Akun")
	fmt.Println("2. Tampilkan Semua Akun")
	fmt.Println("3. Ubah Akun")
	fmt.Println("4. Hapus Akun")
	fmt.Println("5. Cari Akun")
	fmt.Println("6. Urutkan Data")
	fmt.Println("7. Statistik")
	fmt.Println("8. Logout")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func menuUser() {
	fmt.Println("\n===== SELAMAT DATANG DI SECUREPASS =====")
	fmt.Println("1. Register User Baru")
	fmt.Println("2. Login")
	fmt.Println("0. Keluar")
}

func main() {
	var pilihMenuUtama, pilihMenuUser int
	var arrUser tabUsers
	var currentUserIndex int
	var totalUser int

	currentUserIndex = -1
	pilihMenuUser = -1
	totalUser = 0

	for pilihMenuUser != 0 {
		if currentUserIndex == -1 {
			menuUser()
			fmt.Scan(&pilihMenuUser)

			if pilihMenuUser == 1 {
				tambahUser()
			} else if pilihMenuUser == 2 {
				loginUser()
			} else if pilihMenuUser == 0 {
				fmt.Println("Terima kasih sudah menggunakan SecurePass.")
				pilihMenuUser = 0
			} else {
				fmt.Println("Menu tidak valid!")
			}
		} else {
			pilihMenuUtama = -1

			for pilihMenuUtama != 0 && currentUserIndex != -1 {
				menuUtama()
				fmt.Scan(&pilihMenuUtama)

				if pilihMenuUtama == 1 {
					tambahAkun()
				} else if pilihMenuUtama == 2 {
					tampilkanAkun()
				} else if pilihMenuUtama == 3 {
					ubahAkun()
				} else if pilihMenuUtama == 4 {
					hapusAkun()
				} else if pilihMenuUtama == 5 {
					menuCari()
				} else if pilihMenuUtama == 6 {
					menuSort()
				} else if pilihMenuUtama == 7 {
					statistik()
				} else if pilihMenuUtama == 8 {
					fmt.Println("Logout berhasil!")
					currentUserIndex = -1
				} else if pilihMenuUtama == 0 {
					fmt.Println("Terima kasih sudah menggunakan SecurePass.")
					pilihMenuUser = 0
				} else {
					fmt.Println("Menu tidak valid!")
				}	
			}
		}
	}
}