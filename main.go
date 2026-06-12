package main

import "fmt"

const MAX int = 1000

type Waktu struct {
	hari  int
	bulan int
	tahun int
}

type DataTugas struct {
	judul    string
	kategori string
	deadline Waktu
	selesai  bool
}

type ArrTugas [MAX]DataTugas

var daftarTugas ArrTugas
var jumTugas int = 0

func hitungSisaTugas(T ArrTugas, n int) int {
	if n == 0 {
		return 0
	}
	var hitung int = 0
	if T[n-1].selesai == false {
		hitung = 1
	}
	return hitung + hitungSisaTugas(T, n-1)
}

func cekTanggal(h, b, t int) bool {
	if b < 1 || b > 12 || h < 1 || t < 1900 {
		return false
	}
	var batas int = 31
	if b == 4 || b == 6 || b == 9 || b == 11 {
		batas = 30
	} else if b == 2 {
		if (t%400 == 0) || (t%4 == 0 && t%100 != 0) {
			batas = 29
		} else {
			batas = 28
		}
	}
	return h <= batas
}

func tambahTugas(T *ArrTugas, n *int) {
	if *n >= MAX {
		fmt.Println("Gagal: Kapasitas penyimpanan tugas penuh.")
		return
	}
	
	var baru DataTugas
	fmt.Print("Masukkan Judul Tugas (Gunakan_Underscore_Untuk_Spasi): ")
	fmt.Scan(&baru.judul)
	fmt.Print("Masukkan Kategori Tugas: ")
	fmt.Scan(&baru.kategori)

	var h, b, t int
	var valid bool = false

	for valid == false {
		fmt.Print("Masukkan Deadline (Tgl Bln Thn dipisah spasi, cth: 30 5 2026): ")
		fmt.Scan(&h, &b, &t)
		
		if cekTanggal(h, b, t) {
			valid = true
		} else {
			fmt.Println("Error: Tanggal tidak valid. Silakan masukkan ulang.")
		}
	}

	baru.deadline.hari = h
	baru.deadline.bulan = b
	baru.deadline.tahun = t
	baru.selesai = false

	T[*n] = baru
	*n = *n + 1
	fmt.Println("Tugas berhasil ditambahkan.")
}

func tampilTugas(T ArrTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada data tugas yang tercatat.")
		return
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("No | Judul Tugas          | Kategori     | Deadline     | Status")
	fmt.Println("------------------------------------------------------------------")
	
	var i int
	for i = 0; i < n; i++ {
		var statusStr string = "Belum Selesai"
		if T[i].selesai == true {
			statusStr = "Selesai"
		}
		fmt.Printf("%d  | %s | %s | %d-%d-%d | %s\n", i+1, T[i].judul, T[i].kategori, T[i].deadline.hari, T[i].deadline.bulan, T[i].deadline.tahun, statusStr)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Total Tugas Belum Selesai (Rekursif):", hitungSisaTugas(T, n))
}

func editTugas(T *ArrTugas, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data tugas untuk diedit.")
		return
	}
	
	var idx int
	fmt.Print("Masukkan nomor urut tugas yang ingin diedit: ")
	fmt.Scan(&idx)

	if idx <= 0 || idx > n {
		fmt.Println("Error: Nomor urut tidak ditemukan dalam daftar.")
		return
	}
	idx = idx - 1 

	fmt.Println("Atribut apa yang ingin diubah?")
	fmt.Println("1. Tandai Selesai / Belum Selesai")
	fmt.Println("2. Ubah Judul & Kategori")
	
	var pil int
	fmt.Print("Pilihan (1-2): ")
	fmt.Scan(&pil)

	if pil == 1 {
		T[idx].selesai = !T[idx].selesai
		fmt.Println("Status penyelesaian tugas berhasil diperbarui.")
	} else if pil == 2 {
		fmt.Print("Masukkan Judul Baru: ")
		fmt.Scan(&T[idx].judul)
		fmt.Print("Masukkan Kategori Baru: ")
		fmt.Scan(&T[idx].kategori)
		fmt.Println("Data tugas berhasil diubah.")
	} else {
		fmt.Println("Error: Pilihan sub-menu tidak valid.")
	}
}

func hapusTugas(T *ArrTugas, n *int) {
	if *n == 0 {
		fmt.Println("Tidak ada data tugas untuk dihapus.")
		return
	}
	
	var idx int
	fmt.Print("Masukkan nomor urut tugas yang ingin dihapus: ")
	fmt.Scan(&idx)

	if idx <= 0 || idx > *n {
		fmt.Println("Error: Nomor urut tidak ditemukan dalam daftar.")
		return
	}
	idx = idx - 1

	var i int
	for i = idx; i < *n-1; i++ {
		T[i] = T[i+1]
	}
	*n = *n - 1
	fmt.Println("Tugas berhasil dihapus secara permanen.")
}

func deadlineLebihAwal(w1, w2 Waktu) bool {
	if w1.tahun != w2.tahun {
		return w1.tahun < w2.tahun
	}
	if w1.bulan != w2.bulan {
		return w1.bulan < w2.bulan
	}
	return w1.hari < w2.hari
}

func urutDeadline(T *ArrTugas, n int, urutNaik bool) {
	var i, j int
	var temp DataTugas
	var tukar bool

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		tukar = false
		
		if j >= 0 {
			if urutNaik {
				tukar = deadlineLebihAwal(temp.deadline, T[j].deadline)
			} else {
				tukar = deadlineLebihAwal(T[j].deadline, temp.deadline)
			}
		}
		
		for j >= 0 && tukar {
			T[j+1] = T[j]
			j = j - 1
			if j >= 0 {
				if urutNaik {
					tukar = deadlineLebihAwal(temp.deadline, T[j].deadline)
				} else {
					tukar = deadlineLebihAwal(T[j].deadline, temp.deadline)
				}
			}
		}
		T[j+1] = temp
	}
	fmt.Println("Proses pengurutan dengan Insertion Sort selesai.")
}

func urutStatus(T *ArrTugas, n int, urutNaik bool) {
	var i, j, idxMin int
	var valJ, valMin int
	var temp DataTugas

	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			valJ = 0
			valMin = 0
			
			if T[j].selesai == true {
				valJ = 1
			}
			if T[idxMin].selesai == true {
				valMin = 1
			}

			if urutNaik && valJ < valMin {
				idxMin = j
			} else if !urutNaik && valJ > valMin {
				idxMin = j
			}
		}
		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
	fmt.Println("Proses pengurutan dengan Selection Sort selesai.")
}

func cariKategori(T ArrTugas, n int, kataKunci string) {
	var ketemu bool = false
	fmt.Printf("\nHasil Pencarian Kategori [%s]:\n", kataKunci)
	
	var i int
	for i = 0; i < n; i++ {
		if T[i].kategori == kataKunci {
			ketemu = true
			var statusStr string = "Belum Selesai"
			if T[i].selesai == true {
				statusStr = "Selesai"
			}
			fmt.Printf("- %s (DL: %d-%d-%d) [%s]\n", T[i].judul, T[i].deadline.hari, T[i].deadline.bulan, T[i].deadline.tahun, statusStr)
		}
	}
	
	if ketemu == false {
		fmt.Println("Tidak ada tugas yang cocok dengan kategori tersebut.")
	}
}

func cariDeadline(T ArrTugas, n int, h, b, t int) int {
	var kiri int = 0
	var kanan int = n - 1
	var idxKetemu int = -1
	var ketemu bool = false
	
	var target Waktu
	target.hari = h
	target.bulan = b
	target.tahun = t

	for kiri <= kanan && ketemu == false {
		var tengah int = (kiri + kanan) / 2
		if T[tengah].deadline.hari == target.hari && T[tengah].deadline.bulan == target.bulan && T[tengah].deadline.tahun == target.tahun {
			ketemu = true
			idxKetemu = tengah
		} else if deadlineLebihAwal(T[tengah].deadline, target) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idxKetemu
}

func tampilMenu() int {
	var pilihan int
	fmt.Println("\n Sistem Manajemen Tugas Harian  ")
	fmt.Println("=====================================")
	fmt.Println("1. Tambah Tugas Baru")
	fmt.Println("2. Lihat Semua Daftar Tugas")
	fmt.Println("3. Edit Atribut / Status Tugas")
	fmt.Println("4. Hapus Tugas Permanen")
	fmt.Println("5. Cari Tugas")
	fmt.Println("6. Urutkan Tugas")
	fmt.Println("7. Keluar Aplikasi")
	fmt.Println("=====================================")
	
	fmt.Print("Pilih Menu (1-7): ")
	fmt.Scan(&pilihan)
	return pilihan
}

func menuCari(T *ArrTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada data tugas untuk dicari.")
		return
	}

	var subPil int
	fmt.Println("\n--- Sub-Menu Pencarian Data ---")
	fmt.Println("1. Cari Berdasarkan Kategori (Sequential Search)")
	fmt.Println("2. Cari Berdasarkan Tanggal Deadline (Binary Search)")
	fmt.Print("Pilihan: ")
	fmt.Scan(&subPil)
	
	if subPil == 1 {
		var kataKunci string
		fmt.Print("Masukkan Kategori yang dicari: ")
		fmt.Scan(&kataKunci)
		cariKategori(*T, n, kataKunci)
	} else if subPil == 2 {
		fmt.Println("Catatan: Data diurutkan menaik otomatis untuk Binary Search.")
		urutDeadline(T, n, true)
		
		var h, b, t int
		fmt.Print("Masukkan deadline yang dicari (Tgl Bln Thn spasi): ")
		fmt.Scan(&h, &b, &t)
		
		var hasil int = cariDeadline(*T, n, h, b, t)
		if hasil != -1 {
			fmt.Printf("\nData ditemukan di indeks %d: %s [%s]\n", hasil+1, (*T)[hasil].judul, (*T)[hasil].kategori)
		} else {
			fmt.Println("Data tugas dengan tanggal tersebut tidak ditemukan.")
		}
	} else {
		fmt.Println("Error: Pilihan sub-menu pencarian tidak valid.")
	}
}

func menuUrut(T *ArrTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada data tugas untuk diurutkan.")
		return
	}

	var subPil int
	fmt.Println("\n--- Sub-Menu Pengurutan Data ---")
	fmt.Println("1. Urutkan Berdasar Tanggal Deadline (Insertion Sort)")
	fmt.Println("2. Urutkan Berdasar Status Kerja (Selection Sort)")
	fmt.Print("Pilihan Kriteria: ")
	fmt.Scan(&subPil)
	
	if subPil == 1 || subPil == 2 {
		var mode int
		fmt.Println("\nMode Pengurutan:\n1. Ascending (Menaik)\n2. Descending (Menurun)")
		fmt.Print("Pilihan Mode: ")
		fmt.Scan(&mode)
		
		var naik bool = true
		if mode == 2 {
			naik = false
		}

		if subPil == 1 {
			urutDeadline(T, n, naik)
			tampilTugas(*T, n)
		} else if subPil == 2 {
			urutStatus(T, n, naik)
			tampilTugas(*T, n)
		}
	} else {
		fmt.Println("Error: Pilihan kriteria tidak valid.")
	}
}

func main() {
	var pilihan int = 0
	
	for pilihan != 7 {
		pilihan = tampilMenu() 

		if pilihan == 1 {
			tambahTugas(&daftarTugas, &jumTugas)
		} else if pilihan == 2 {
			tampilTugas(daftarTugas, jumTugas)
		} else if pilihan == 3 {
			editTugas(&daftarTugas, jumTugas)
		} else if pilihan == 4 {
			hapusTugas(&daftarTugas, &jumTugas)
		} else if pilihan == 5 {
			menuCari(&daftarTugas, jumTugas)
		} else if pilihan == 6 {
			menuUrut(&daftarTugas, jumTugas)
		} else if pilihan == 7 {
			fmt.Println("Terima kasih telah menggunakan aplikasi. Selesai.")
		} else {
			fmt.Println("Error: Pilihan tidak valid, mohon masukkan angka 1-7.")
		}
	}
}
