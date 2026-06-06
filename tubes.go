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
	/* mengembalikan string input dari pengguna berdasarkan pesan yang dicetak */
	var x string
	fmt.Print(pesan)
	fmt.Scan(&x)
	return x
}

func inputTanggal() string {
	/* mengembalikan string tanggal dan waktu saat ini dengan format YYYY-MM-DD HH:MM:SS */
	var tgl string

	tgl = time.Now().Format("2006-01-02 15:04:05")

	return tgl
}

func cekKekuatanPassword(pass string) string {
	/* mengembalikan string kategori kekuatan password berupa "LEMAH", "SEDANG", atau "KUAT" berdasarkan panjang dan kombinasi karakternya */
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
	/* 
		I.S. terdefinisi array user a dan n jumlah total user.
	  F.S. apabila n < userMax dan username belum digunakan, data user baru ditambahkan dan n bertambah 1. Apabila user penuh atau duplikat, maka akan menampilkan pesan error.
	*/
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
	/*
		I.S. terdefinisi array user a, jumlah total user n, dan index user yang sedang login.
		F.S. index user yang sedang login berisi angka index dari user yang cocok. Apabila n == 0, program memanggil prosedur tambahUser.
	*/
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
	/*
		I.S. terdefinisi array user a dan index user yang sedang login.
		F.S. apabila kapasitas belum penuh, data akun baru ditambahkan ke dalam array kumpulanAkun milik user tersebut dan jumlahAkun bertambah 1
	*/
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
	/*
		I.S. terdefinisi array user a dan index user yang sedang login.
		F.S. menampilkan seluruh data dari kumpulanAkun milik user.
	*/
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
	/*
		I.S. terdefinisi array user a dan index user yang sedang login.
		F.S. data akun pada nomor indeks yang dipilih pengguna diperbarui dengan data layanan, email, password, dan waktu update yang baru.
	*/
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
	/*
		I.S. terdefinisi array user a dan index user yang sedang login.
		F.S. data akun pada nomor indeks yang dipilih dihapus, index di belakangnya digeser maju, dan jumlahAkun berkurang 1.
	*/
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

func sequentialSearch(a tabUsers, userIndex int, email string, result *[akunMax]int, resultTotal *int) {
	/*
		I.S. terdefinisi array user a, index user, keyword email/username yang dicari, array integer pencarian result, dan jumlah total hasil pencarian resultTotal.
		F.S. array result berisi indeks dari semua akun dengan email/username yang cocok, dan resultTotal berisi jumlah total data yang ditemukan.
	*/
	var i int
	var n int 
	
	n = a[userIndex].jumlahAkun
	*resultTotal = 0

	for i = 0; i < n; i++ {
		if a[userIndex].kumpulanAkun[i].email == email {
			result[*resultTotal] = i
			*resultTotal++
		}
	}
}

func binarySearch(a tabUsers, userIndex int, layanan string, result *[akunMax]int, resultTotal *int) {
	/*
		I.S. terdefinisi array user a, index user, keyword email/username yang dicari, array integer pencarian result, dan jumlah total hasil pencarian resultTotal.
				 Data kumpulanAkun sudah terurut membesar (ascending) berdasarkan layanan.
		F.S. array result berisi indeks dari semua akun dengan nama layanan yang cocok, dan resultTotal berisi jumlah total data yang ditemukan.
	*/
	var kiri, kanan, tengah, i, batasKiri, batasKanan int
	var n int 
	var found bool

	found = false
	n = a[userIndex].jumlahAkun
	kiri = 0
	kanan = n - 1
	*resultTotal = 0

	for kiri <= kanan && !found {
		tengah = (kiri + kanan) / 2

		if a[userIndex].kumpulanAkun[tengah].layanan == layanan {
			found = true
		} else if a[userIndex].kumpulanAkun[tengah].layanan < layanan {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if found {
		batasKiri = tengah
		for batasKiri > 0 && a[userIndex].kumpulanAkun[batasKiri-1].layanan == layanan {
			batasKiri--
		}

		batasKanan = tengah
		for batasKanan < n-1 && a[userIndex].kumpulanAkun[batasKanan+1].layanan == layanan {
			batasKanan++
		}

		for i = batasKiri; i <= batasKanan; i++ {
			result[*resultTotal] = i
			*resultTotal++
		}
	}
}

func menuCari(a *tabUsers, userIndex int) {
	/*
		I.S. terdefinisi pointer user a dan index user yang sedang login.
		F.S. menampilkan menu pencarian, memanggil subprogram pencarian yang sesuai dengan input yang dipilih, dan menampilkan hasil pencarian akun.
	*/
	var i, pilih int
	var pos int
	var layanan string
	var results [akunMax]int
	var resultTotal int

	fmt.Println("\n=== MENU PENCARIAN ===")
	fmt.Println("1. Cari berdasarkan email/username")
	fmt.Println("2. Cari berdasarkan nama layanan.")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	resultTotal = 0

	if pilih == 1 {
		layanan = inputString("Masukkan username/email yang dicari: ")
		sequentialSearch(*a, userIndex, layanan, &results, &resultTotal)
	} else if pilih == 2 {
		layanan = inputString("Masukkan nama layanan yang dicari: ")
		selectionSortNama(a, userIndex, 1)
		binarySearch(*a, userIndex, layanan, &results, &resultTotal)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}

	if resultTotal > 0 {
		fmt.Println("\nData ditemukan! Jumlah:", resultTotal)
		for i = 0; i < resultTotal; i++ {
			pos = results[i] 
			
			fmt.Println("----------------------------")
			fmt.Println("Layanan :", a[userIndex].kumpulanAkun[pos].layanan)
			fmt.Println("Email   :", a[userIndex].kumpulanAkun[pos].email)
			fmt.Println("Pass    :", a[userIndex].kumpulanAkun[pos].password)
			fmt.Println("Update  :", a[userIndex].kumpulanAkun[pos].lastUpdate)
			fmt.Println("Kekuatan:", cekKekuatanPassword(a[userIndex].kumpulanAkun[pos].password))
		}
		fmt.Println("----------------------------")
	} else if pilih == 1 || pilih == 2 {
		fmt.Println("\nData tidak ditemukan.") 
	}
}

func selectionSortNama(a *tabUsers, userIndex int, sortOption int) {
	/*
		I.S. terdefinisi array user a, index user yang sedang login, dan pilihan arah urutan sortOption dalam integer.
		F.S. data pada kumpulanAkun terurut berdasarkan layanan secara membesar jika sortOption = 1 atau mengecil jika sortOption = 2.
	*/
	var i, j, min int
	var n int = a[userIndex].jumlahAkun
	var temp Akun

	if sortOption == 1 /* Ascending*/ {
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
	} else if sortOption == 2 /*Descending*/ {
		for i = 0; i < n-1; i++ {
			min = i
			for j = i + 1; j < n; j++ {
				if a[userIndex].kumpulanAkun[j].layanan > a[userIndex].kumpulanAkun[min].layanan {
					min = j
				}
			}

			temp = a[userIndex].kumpulanAkun[i]
			a[userIndex].kumpulanAkun[i] = a[userIndex].kumpulanAkun[min]
			a[userIndex].kumpulanAkun[min] = temp
		}
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama layanan (A-Z).")
}

func insertionSortTanggal(a *tabUsers, userIndex int, sortOption int) {
	/*
		I.S. terdefinisi array user a, index user yang sedang login, dan pilihan arah urutan sortOption dalam integer.
		F.S. data pada kumpulanAkun terurut berdasarkan lastUpdate secara membesar jika sortOption = 1 atau mengecil jika sortOption = 2.
	*/
	var i, j int
	var n int = a[userIndex].jumlahAkun
	var key Akun

	if sortOption == 1 /*Ascending*/ {
		for i = 1; i < n; i++ {
			key = a[userIndex].kumpulanAkun[i]
			j = i - 1

			for j >= 0 && a[userIndex].kumpulanAkun[j].lastUpdate > key.lastUpdate {
				a[userIndex].kumpulanAkun[j+1] = a[userIndex].kumpulanAkun[j]
				j--
			}

			a[userIndex].kumpulanAkun[j+1] = key
		}
	} else if sortOption == 2 /*Descending*/ {
		for i = 1; i < n; i++ {
			key = a[userIndex].kumpulanAkun[i]
			j = i - 1

			for j >= 0 && a[userIndex].kumpulanAkun[j].lastUpdate < key.lastUpdate {
				a[userIndex].kumpulanAkun[j+1] = a[userIndex].kumpulanAkun[j]
				j--
			}

			a[userIndex].kumpulanAkun[j+1] = key
		}
	}

	

	fmt.Println("Data berhasil diurutkan berdasarkan tanggal update.")
}

func menuSort(a *tabUsers, userIndex int) {
	/*
		I.S. terdefinisi array user a, index user yang sedang login, dan pilihan arah urutan sortOption dalam integer.
		F.S. menampilkan opsi menu sorting, memanggil subprogram sorting yang sesuai dengan input, lalu menampilkan seluruh akun yang sudah terurut.
	*/
	var pilih, pilihOrder int

	fmt.Println("\n=== MENU SORTING ===")
	fmt.Println("1. Urutkan berdasarkan Nama layanan")
	fmt.Println("2. Urutkan berdasarkan Tanggal update")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 /* Selection Sort (Nama Layanan) */ {
		fmt.Println("1. Urutkan secara ascending (A-Z)")
		fmt.Println("2. Urutkan secara descending (Z-A)")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihOrder)
		if pilihOrder == 1 || pilihOrder == 2 {
			selectionSortNama(a, userIndex, pilihOrder)
			fmt.Println("Data berhasil diurutkan berdasarkan nama layanan.")
			tampilkanAkun(*a, userIndex)
		} else {
			fmt.Println("Pilihan arah tidak valid!")
		}
	} else if pilih == 2 /* Insertion Sort (Tanggal Update) */ {
		fmt.Println("1. Urutkan secara ascending (Terlama - Terbaru)")
		fmt.Println("2. Urutkan secara descending (Terbaru - Terlama)")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihOrder)
		if pilihOrder == 1 || pilihOrder == 2 {
			insertionSortTanggal(a, userIndex, pilihOrder)
			fmt.Println("Data berhasil diurutkan berdasarkan tanggal update.")
			tampilkanAkun(*a, userIndex)
		} else {
			fmt.Println("Pilihan arah tidak valid!")
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

func statistik(a tabUsers, userIndex int) {
	/*
		I.S. terdefinisi array user a dan index user yang sedang login.
		F.S. menampilkan statistik total akun tersimpan serta jumlah akun dengan masing-masing kategori password LEMAH, SEDANG, dan KUAT.
	*/
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
	/*
		I.S. terdefinisi string username dari pengguna yang sedang login.
		F.S. menampilkan interface opsi menu utama dari program SecurePass.
	*/
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
	/*
		I.S. 
		F.S. F.S. menampilkan interface opsi menu registrasi dan login user.
	*/
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
					menuCari(&arrUser, currentUserIndex)
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