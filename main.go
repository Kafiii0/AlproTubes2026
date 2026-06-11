package main

import "fmt"

const NMAX int = 1000

type Waktu struct {
	tgl int
	bln int
	thn int
}

type Tugas struct {
	judul     string
	kategori  string
	deadline  Waktu
	isSelesai bool
}

type TabTugas [NMAX]Tugas

var daftarTugas TabTugas
var nTugas int = 0

func hitungTugasBelumSelesaiRekursif(T TabTugas, n int) int {
	if n == 0 {
		return 0
	}
	var hitungSaatIni int = 0
	if !T[n-1].isSelesai {
		hitungSaatIni = 1
	}
	return hitungSaatIni + hitungTugasBelumSelesaiRekursif(T, n-1)
}

func isTanggalValid(tgl, bln, thn int) bool {
	if bln < 1 || bln > 12 || tgl < 1 || thn < 1900 {
		return false
	}
	var batasHari int = 31
	if bln == 4 || bln == 6 || bln == 9 || bln == 11 {
		batasHari = 30
	} else if bln == 2 {
		if (thn%400 == 0) || (thn%4 == 0 && thn%100 != 0) {
			batasHari = 29
		} else {
			batasHari = 28
		}
	}
	return tgl <= batasHari
}

func tambahTugas(T *TabTugas, n *int) {
	if *n >= NMAX {
		fmt.Println("Gagal: Kapasitas penyimpanan tugas penuh!")
		return
	}
	
	var t Tugas
	fmt.Print("Masukkan Judul Tugas (Gunakan_Underscore_Untuk_Spasi): ")
	fmt.Scan(&t.judul)
	fmt.Print("Masukkan Kategori Tugas: ")
	fmt.Scan(&t.kategori)

	var tgl, bln, thn int
	var valid bool = false

	for !valid {
		fmt.Print("Masukkan Deadline (Tgl Bln Thn dipisah spasi, cth: 30 5 2026): ")
		fmt.Scan(&tgl, &bln, &thn)
		
		if isTanggalValid(tgl, bln, thn) {
			valid = true
		} else {
			fmt.Println("Error: Tanggal tidak valid! Silakan masukkan ulang.")
		}
	}

	t.deadline.tgl = tgl
	t.deadline.bln = bln
	t.deadline.thn = thn
	t.isSelesai = false

	T[*n] = t
	*n = *n + 1
	fmt.Println("Tugas berhasil ditambahkan!")
}

func tampilSemuaTugas(T TabTugas, n int) {
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
		if T[i].isSelesai {
			statusStr = "Selesai"
		}
		fmt.Printf("%d  | %s | %s | %d-%d-%d | %s\n", i+1, T[i].judul, T[i].kategori, T[i].deadline.tgl, T[i].deadline.bln, T[i].deadline.thn, statusStr)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Total Tugas Belum Selesai (Rekursif):", hitungTugasBelumSelesaiRekursif(T, n))
}

func editTugas(T *TabTugas, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data tugas untuk diedit.")
		return
	}
	
	var idx int
	fmt.Print("Masukkan nomor urut tugas yang ingin diedit: ")
	fmt.Scan(&idx)

	if idx <= 0 || idx > n {
		fmt.Println("Error: Nomor urut tidak ditemukan dalam daftar!")
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
		T[idx].isSelesai = !T[idx].isSelesai
		fmt.Println("Status penyelesaian tugas berhasil diperbarui!")
	} else if pil == 2 {
		fmt.Print("Masukkan Judul Baru: ")
		fmt.Scan(&T[idx].judul)
		fmt.Print("Masukkan Kategori Baru: ")
		fmt.Scan(&T[idx].kategori)
		fmt.Println("Data tugas berhasil diubah!")
	} else {
		fmt.Println("Error: Pilihan sub-menu tidak valid.")
	}
}

func hapusTugas(T *TabTugas, n *int) {
	if *n == 0 {
		fmt.Println("Tidak ada data tugas untuk dihapus.")
		return
	}
	
	var idx int
	fmt.Print("Masukkan nomor urut tugas yang ingin dihapus: ")
	fmt.Scan(&idx)

	if idx <= 0 || idx > *n {
		fmt.Println("Error: Nomor urut tidak ditemukan dalam daftar!")
		return
	}
	idx = idx - 1

	var i int
	for i = idx; i < *n-1; i++ {
		T[i] = T[i+1]
	}
	*n = *n - 1
	fmt.Println("Tugas berhasil dihapus secara permanen!")
}

func isDeadlineLebihKecil(w1, w2 Waktu) bool {
	if w1.thn != w2.thn {
		return w1.thn < w2.thn
	}
	if w1.bln != w2.bln {
		return w1.bln < w2.bln
	}
	return w1.tgl < w2.tgl
}

func urutBerdasarkanDeadline(T *TabTugas, n int, ascending bool) {
	var i, j int
	var key Tugas
	var kondisi bool

	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		kondisi = false
		
		if j >= 0 {
			if ascending {
				kondisi = isDeadlineLebihKecil(key.deadline, T[j].deadline)
			} else {
				kondisi = isDeadlineLebihKecil(T[j].deadline, key.deadline)
			}
		}
		
		for j >= 0 && kondisi {
			T[j+1] = T[j]
			j = j - 1
			if j >= 0 {
				if ascending {
					kondisi = isDeadlineLebihKecil(key.deadline, T[j].deadline)
				} else {
					kondisi = isDeadlineLebihKecil(T[j].deadline, key.deadline)
				}
			}
		}
		T[j+1] = key
	}
	fmt.Println("Proses pengurutan dengan Insertion Sort selesai.")
}

func urutBerdasarkanStatus(T *TabTugas, n int, ascending bool) {
	var i, j, idxEkstrem int
	var valJ, valEkstrem int
	var temp Tugas

	for i = 0; i < n-1; i++ {
		idxEkstrem = i
		for j = i + 1; j < n; j++ {
			valJ = 0
			valEkstrem = 0
			
			if T[j].isSelesai {
				valJ = 1
			}
			if T[idxEkstrem].isSelesai {
				valEkstrem = 1
			}

			if ascending && valJ < valEkstrem {
				idxEkstrem = j
			} else if !ascending && valJ > valEkstrem {
				idxEkstrem = j
			}
		}
		temp = T[i]
		T[i] = T[idxEkstrem]
		T[idxEkstrem] = temp
	}
	fmt.Println("Proses pengurutan dengan Selection Sort selesai.")
}

func cariBerdasarkanKategori(T TabTugas, n int, keyword string) {
	var found bool = false
	fmt.Printf("\nHasil Pencarian Kategori [%s]:\n", keyword)
	
	var i int
	for i = 0; i < n; i++ {
		if T[i].kategori == keyword {
			found = true
			var statusStr string = "Belum Selesai"
			if T[i].isSelesai {
				statusStr = "Selesai"
			}
			fmt.Printf("- %s (DL: %d-%d-%d) [%s]\n", T[i].judul, T[i].deadline.tgl, T[i].deadline.bln, T[i].deadline.thn, statusStr)
		}
	}
	
	if !found {
		fmt.Println("Tidak ada tugas yang cocok dengan kategori tersebut.")
	}
}

func cariBerdasarkanDeadline(T TabTugas, n int, tgl, bln, thn int) int {
	var left int = 0
	var right int = n - 1
	var foundIdx int = -1
	var found bool = false
	
	var target Waktu
	target.tgl = tgl
	target.bln = bln
	target.thn = thn

	for left <= right && !found {
		var mid int = (left + right) / 2
		if T[mid].deadline.tgl == target.tgl && T[mid].deadline.bln == target.bln && T[mid].deadline.thn == target.thn {
			found = true
			foundIdx = mid
		} else if isDeadlineLebihKecil(T[mid].deadline, target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return foundIdx
}


func menuUtama() {
	var pilihan int = 0
	for pilihan != 7 {
		fmt.Println("\n=====================================")
		fmt.Println("    APLIKASI MANAJEMEN TUGAS HARIAN  ")
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

		if pilihan == 1 {
			tambahTugas(&daftarTugas, &nTugas)
		} else if pilihan == 2 {
			tampilSemuaTugas(daftarTugas, nTugas)
		} else if pilihan == 3 {
			editTugas(&daftarTugas, nTugas)
		} else if pilihan == 4 {
			hapusTugas(&daftarTugas, &nTugas)
		} else if pilihan == 5 {
			var subPil int
			fmt.Println("\n--- Sub-Menu Pencarian Data ---")
			fmt.Println("1. Cari Berdasarkan Kategori (Sequential Search)")
			fmt.Println("2. Cari Berdasarkan Tanggal Deadline (Binary Search)")
			fmt.Print("Pilihan: ")
			fmt.Scan(&subPil)
			
			if subPil == 1 {
				var key string
				fmt.Print("Masukkan Kategori yang dicari: ")
				fmt.Scan(&key)
				cariBerdasarkanKategori(daftarTugas, nTugas, key)
			} else if subPil == 2 {
				fmt.Println("Catatan: Data diurutkan menaik otomatis untuk Binary Search.")
				urutBerdasarkanDeadline(&daftarTugas, nTugas, true)
				
				var t, b, th int
				fmt.Print("Masukkan deadline yang dicari (Tgl Bln Thn spasi): ")
				fmt.Scan(&t, &b, &th)
				
				var resIdx int = cariBerdasarkanDeadline(daftarTugas, nTugas, t, b, th)
				if resIdx != -1 {
					fmt.Printf("\nData ditemukan di indeks %d: %s [%s]\n", resIdx+1, daftarTugas[resIdx].judul, daftarTugas[resIdx].kategori)
				} else {
					fmt.Println("Data tugas dengan tanggal tersebut tidak ditemukan.")
				}
			} else {
				fmt.Println("Error: Pilihan sub-menu pencarian tidak valid.")
			}
		} else if pilihan == 6 {
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
				
				var isAsc bool = true
				if mode == 2 {
					isAsc = false
				}

				if subPil == 1 {
					urutBerdasarkanDeadline(&daftarTugas, nTugas, isAsc)
					tampilSemuaTugas(daftarTugas, nTugas)
				} else if subPil == 2 {
					urutBerdasarkanStatus(&daftarTugas, nTugas, isAsc)
					tampilSemuaTugas(daftarTugas, nTugas)
				}
			} else {
				fmt.Println("Error: Pilihan kriteria tidak valid.")
			}
		} else if pilihan == 7 {
			fmt.Println("Terima kasih telah menggunakan aplikasi! Selesai.")
		} else {
			fmt.Println("Error: Pilihan tidak valid, mohon masukkan angka 1-7.")
		}
	}
}


func main() {
	menuUtama()
}
