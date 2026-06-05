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

func tambahUser(a *tabUsers, n *int) {
	var i int
	var newUsername string
	var isDuplicate bool

	isDuplicate = false

	if *n < userMax {
		fmt.Println("\n=== TAMBAH USER ===")

		newUsername = inputString("Username      :")

		for i = 0; i < *n; i++ {
			if a[i].username == newUsername {
				isDuplicate = true
			}
		}

		if isDuplicate {
			fmt.Println("Gagal! Username sudah digunakan. Silakan coba yang lain.")
		} else {
			a[*n].username = newUsername
			a[*n].password = inputString("Password      :")

			*n++

			fmt.Println("User berhasil ditambahkan!")
		}
	} else {
		fmt.Println("Data penuh! Tidak bisa menambah user.")
	}
}

func loginUser(a *tabUsers, n *int, userIndex *int) {
	var inputUsername, inputPassword string
	var i int
	var isFound bool

	isFound = false

	if *n == 0 {
		fmt.Println("\n=== LOGIN USER ===")
		fmt.Println("Error! Tidak ada User yang terdaftar!")
		fmt.Println("Silahkan lakukan Register terlebih dahulu!")
		tambahUser(a, n)
	} else {
		fmt.Println("\n=== LOGIN USER ===")
		inputUsername = inputString("Username    :")
		inputPassword = inputString("Password    :")

		for i = 0; i < *n; i++ {
			if a[i].username == inputUsername && a[i].password == inputPassword {
				*userIndex = i
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

func tambahAkun(a *tabUsers, userIndex int) {
	var jumlahAkun int 
	
	jumlahAkun = a[userIndex].jumlahAkun

	if jumlahAkun < akunMax {
		fmt.Println("\n=== TAMBAH AKUN ===")

		a[userIndex].kumpulanAkun[jumlahAkun].layanan = inputString("Nama Layanan      : ")
		a[userIndex].kumpulanAkun[jumlahAkun].email = inputString("Email/Username    : ")
		a[userIndex].kumpulanAkun[jumlahAkun].password = inputString("Password          : ")
		a[userIndex].kumpulanAkun[jumlahAkun].lastUpdate = inputTanggal()

		a[userIndex].jumlahAkun++
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Data penuh! Tidak bisa menambah akun.")
	}
}

func tampilkanAkun(a tabUsers, userIndex int) {
	var i int
	var n int = a[userIndex].jumlahAkun
	fmt.Println("\n=== DAFTAR AKUN ===")

	if n == 0 {
		fmt.Println("Belum ada data akun.")
	} else {
		for i = 0; i < n; i++ {
			fmt.Println("----------------------------")
			fmt.Println("No          :", i+1)
			fmt.Println("Layanan     :", a[userIndex].kumpulanAkun[i].layanan)
			fmt.Println("Email       :", a[userIndex].kumpulanAkun[i].email)
			fmt.Println("Last Update :", a[userIndex].kumpulanAkun[i].lastUpdate)
		}
		fmt.Println("----------------------------")
	}
}

func ubahAkun(a *tabUsers, userIndex int) {
	var idx, i int
	var n int = a[userIndex].jumlahAkun
	fmt.Println("\n=== UBAH AKUN ===")
	tampilkanAkun(*a, userIndex)

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa diubah.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin diubah: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			i = idx - 1

			fmt.Println("\nData lama:")
			fmt.Println("Layanan :", a[userIndex].kumpulanAkun[i].layanan)
			fmt.Println("Email   :", a[userIndex].kumpulanAkun[i].email)
			fmt.Println("Pass    :", a[userIndex].kumpulanAkun[i].password)
			fmt.Println("Update  :", a[userIndex].kumpulanAkun[i].lastUpdate)

			a[userIndex].kumpulanAkun[i].layanan = inputString("Layanan baru   : ")
			a[userIndex].kumpulanAkun[i].email = inputString("Email baru     : ")
			a[userIndex].kumpulanAkun[i].password = inputString("Password baru  : ")
			a[userIndex].kumpulanAkun[i].lastUpdate = inputTanggal()

			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

func hapusAkun(a *tabUsers, userIndex int) {
	var idx, i, pos int
	var n int = a[userIndex].jumlahAkun
	fmt.Println("\n=== HAPUS AKUN ===")
	tampilkanAkun(*a, userIndex)

	if n == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa dihapus.")
	} else {
		fmt.Print("Masukkan nomor akun yang ingin dihapus: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= n {
			pos = idx - 1

			for i = pos; i < n-1; i++ {
				a[userIndex].kumpulanAkun[i] = a[userIndex].kumpulanAkun[i+1]
			}

			a[userIndex].jumlahAkun--
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("Nomor tidak valid!")
		}
	}
}

// searching 
// Sequential Search (layanan)
func sequentialSearch(a tabUsers, userIndex int, layanan string) int {
	var pos, i int
	pos = -1
	var n int = a[userIndex].jumlahAkun

	for i = 0; i < n; i++ {
		if a[userIndex].kumpulanAkun[i].layanan == layanan {
			pos = i
		}
	}
	return pos
}

// Binary Search (data harus sudah diurutkan alfabet) (username atau email)
func binarySearch(a tabUsers, userIndex int, email string) int {
	var kiri, kanan, pos, tengah int
	var n int = a[userIndex].jumlahAkun

	kiri = 0
	kanan = n - 1
	pos = -1

	for kiri <= kanan && pos == -1 {
		tengah = (kiri + kanan) / 2

		if a[userIndex].kumpulanAkun[tengah].email == email {
			pos = tengah
		} else if a[userIndex].kumpulanAkun[tengah].email < email {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return pos
}

func menuCari(a tabUsers, userIndex int) {
	var pilih, pos int
	var layanan string

	fmt.Println("\n=== MENU PENCARIAN ===")
	fmt.Println("1. Cari berdasarkan layanan")
	fmt.Println("2. Cari berdasarkan email")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	layanan = inputString("Masukkan username/email yang dicari: ")

	if pilih == 1 {
		pos = sequentialSearch(a, userIndex, layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", a[userIndex].kumpulanAkun[pos].layanan)
			fmt.Println("Email   :", a[userIndex].kumpulanAkun[pos].email)
			fmt.Println("Pass    :", a[userIndex].kumpulanAkun[pos].password)
			fmt.Println("Update  :", a[userIndex].kumpulanAkun[pos].lastUpdate)
			fmt.Println("Kekuatan:", cekKekuatanPassword(a[userIndex].kumpulanAkun[pos].password))
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	} else if pilih == 2 {
		pos = binarySearch(a, userIndex, layanan)
		if pos != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Println("Layanan :", a[userIndex].kumpulanAkun[pos].layanan)
			fmt.Println("Email   :", a[userIndex].kumpulanAkun[pos].email)
			fmt.Println("Pass    :", a[userIndex].kumpulanAkun[pos].password)
			fmt.Println("Update  :", a[userIndex].kumpulanAkun[pos].lastUpdate)
			fmt.Println("Kekuatan:", cekKekuatanPassword(a[userIndex].kumpulanAkun[pos].password))
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

//SORTING 
// Selection Sort berdasarkan nama layanan alfabet
func selectionSortNama(a *tabUsers, userIndex int) {
	var i, j, min int
	var n int = a[userIndex].jumlahAkun
	var temp Akun

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if a[userIndex].kumpulanAkun[j].layanan < a[userIndex].kumpulanAkun[min].layanan {
				min = j
			}
		}

		temp = a[userIndex].kumpulanAkun[i]
		a[userIndex].kumpulanAkun[i] = a[userIndex].kumpulanAkun[min]
		a[userIndex].kumpulanAkun[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama layanan (A-Z).")
}

// Insertion Sort berdasarkan tanggal update
func insertionSortTanggal(a *tabUsers, userIndex int) {
	var i, j int
	var n int = a[userIndex].jumlahAkun
	var key Akun

	for i = 1; i < n; i++ {
		key = a[userIndex].kumpulanAkun[i]
		j = i - 1

		for j >= 0 && a[userIndex].kumpulanAkun[j].lastUpdate > key.lastUpdate {
			a[userIndex].kumpulanAkun[j+1] = a[userIndex].kumpulanAkun[j]
			j--
		}

		a[userIndex].kumpulanAkun[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tanggal update.")
}

func menuSort(a *tabUsers, userIndex int) {
	var pilih int

	fmt.Println("\n=== MENU SORTING ===")
	fmt.Println("1. Selection Sort (Nama layanan A-Z)")
	fmt.Println("2. Insertion Sort (Tanggal update)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortNama(a, userIndex)
	} else if pilih == 2 {
		insertionSortTanggal(a, userIndex)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

func statistik(a tabUsers, userIndex int) {
	var lemah, kuat, sedang, i int
	var k string
	var n int = a[userIndex].jumlahAkun

	fmt.Println("\n=== STATISTIK SECUREPASS ===")
	fmt.Println("Total akun tersimpan:", n)

	lemah = 0
	sedang = 0
	kuat = 0

	for i = 0; i < n; i++ {
		k = cekKekuatanPassword(a[userIndex].kumpulanAkun[i].password)

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

func menuUtama(username string) {
	fmt.Println("\n===== SECUREPASS MENU =====")
	fmt.Printf("===== Selamat Datang, %s! ====\n", username)
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
	fmt.Print("Pilih: ")
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
				tambahUser(&arrUser, &totalUser)
			} else if pilihMenuUser == 2 {
				loginUser(&arrUser, &totalUser, &currentUserIndex)
			} else if pilihMenuUser == 0 {
				fmt.Println("Terima kasih sudah menggunakan SecurePass.")
				pilihMenuUser = 0
			} else {
				fmt.Println("Menu tidak valid!")
			}
		} else {
			pilihMenuUtama = -1

			for pilihMenuUtama != 0 && currentUserIndex != -1 {
				menuUtama(arrUser[currentUserIndex].username)
				fmt.Print("Pilih: ")
				fmt.Scan(&pilihMenuUtama)

				if pilihMenuUtama == 1 {
					tambahAkun(&arrUser, currentUserIndex)
				} else if pilihMenuUtama == 2 {
					tampilkanAkun(arrUser, currentUserIndex)
				} else if pilihMenuUtama == 3 {
					ubahAkun(&arrUser, currentUserIndex)
				} else if pilihMenuUtama == 4 {
					hapusAkun(&arrUser, currentUserIndex)
				} else if pilihMenuUtama == 5 {
					menuCari(arrUser, currentUserIndex)
				} else if pilihMenuUtama == 6 {
					menuSort(&arrUser, currentUserIndex)
				} else if pilihMenuUtama == 7 {
					statistik(arrUser, currentUserIndex)
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