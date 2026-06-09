package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
var reader = bufio.NewReader(os.Stdin)

func inputStringWajib(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input != "" {
			return input
		}
		fmt.Println("[!] Error: Input tidak boleh kosong! Silakan coba lagi.")
	}
}

func inputIntWajib(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "" {
			fmt.Println("[!] Error: Input tidak boleh kosong! Harus berupa angka.")
			continue
		}

		val, err := strconv.Atoi(input)
		if err == nil {
			return val 
		}
		fmt.Println("[!] Error: Input tidak valid! Harap masukkan huruf/karakter berupa angka.")
	}
}

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
		fmt.Println("[!] Gagal: Kapasitas penyimpanan tugas penuh!")
		return
	}
	var t Tugas
	t.judul = inputStringWajib("Masukkan Judul Tugas: ")
	t.kategori = inputStringWajib("Masukkan Kategori Tugas: ")

	var tgl, bln, thn int
	valid := false

	for !valid {
		inputDL := inputStringWajib("Masukkan Deadline (Tgl Bln Thn, contoh: 30 5 2026): ")
		parts := strings.Split(inputDL, " ")
		
		if len(parts) == 3 {
			var errTgl, errBln, errThn error
			tgl, errTgl = strconv.Atoi(parts[0])
			bln, errBln = strconv.Atoi(parts[1])
			thn, errThn = strconv.Atoi(parts[2])
			
			if errTgl != nil || errBln != nil || errThn != nil {
				fmt.Println("[!] Error: Format harus berupa angka! (contoh: 30 5 2026)")
			} else if isTanggalValid(tgl, bln, thn) {
				valid = true
			} else {
				fmt.Println("[!] Error: Tanggal tidak masuk akal dalam kalender! Silakan masukkan ulang.")
			}
		} else {
			fmt.Println("[!] Error: Format salah! Gunakan spasi untuk memisahkan tanggal, bulan, dan tahun.")
		}
	}

	t.deadline = Waktu{tgl: tgl, bln: bln, thn: thn}
	t.isSelesai = false

	T[*n] = t
	*n = *n + 1
	fmt.Println("-> Tugas berhasil ditambahkan!")
}

func tampilSemuaTugas(T TabTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada data tugas yang tercatat.")
		return
	}
	fmt.Println("\n=========================================================================")
	fmt.Println("No | Judul Tugas          | Kategori     | Deadline     | Status")
	fmt.Println("=========================================================================")
	for i := 0; i < n; i++ {
		var statusStr string = "Belum Selesai"
		if T[i].isSelesai {
			statusStr = "Selesai"
		}
		fmt.Println(i+1, " |", T[i].judul, "|", T[i].kategori, "|", T[i].deadline.tgl, "-", T[i].deadline.bln, "-", T[i].deadline.thn, "|", statusStr)
	}
	fmt.Println("=========================================================================")
	fmt.Println("Total Tugas Belum Selesai (Hitungan Rekursif):", hitungTugasBelumSelesaiRekursif(T, n))
}

func editTugas(T *TabTugas, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data tugas untuk diedit.")
		return
	}
	idx := inputIntWajib("Masukkan nomor urut tugas yang ingin diedit: ")

	if idx <= 0 || idx > n {
		fmt.Println("[!] Error: Nomor urut tidak ditemukan dalam daftar!")
		return
	}
	idx = idx - 1 

	fmt.Println("\nAtribut apa yang ingin diubah?")
	fmt.Println("1. Tandai Selesai / Belum Selesai")
	fmt.Println("2. Ubah Judul & Kategori")
	pil := inputIntWajib("Pilihan (1-2): ")

	if pil == 1 {
		T[idx].isSelesai = !T[idx].isSelesai
		fmt.Println("-> Status penyelesaian tugas berhasil diperbarui!")
	} else if pil == 2 {
		T[idx].judul = inputStringWajib("Masukkan Judul Baru: ")
		T[idx].kategori = inputStringWajib("Masukkan Kategori Baru: ")
		fmt.Println("-> Data tugas berhasil diubah!")
	} else {
		fmt.Println("[!] Error: Pilihan sub-menu tidak valid.")
	}
}

func hapusTugas(T *TabTugas, n *int) {
	if *n == 0 {
		fmt.Println("Tidak ada data tugas untuk dihapus.")
		return
	}
	idx := inputIntWajib("Masukkan nomor urut tugas yang ingin dihapus: ")

	if idx <= 0 || idx > *n {
		fmt.Println("[!] Error: Nomor urut tidak ditemukan dalam daftar!")
		return
	}
	idx = idx - 1

	for i := idx; i < *n-1; i++ {
		T[i] = T[i+1]
	}
	*n = *n - 1
	fmt.Println("-> Tugas berhasil dihapus secara permanen!")
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
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1
		var kondisi bool = false
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
	fmt.Println("-> Proses pengurutan dengan Insertion Sort selesai.")
}

func urutBerdasarkanStatus(T *TabTugas, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		idxEkstrem := i
		for j := i + 1; j < n; j++ {
			var valJ, valEkstrem int = 0, 0
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
		temp := T[i]
		T[i] = T[idxEkstrem]
		T[idxEkstrem] = temp
	}
	fmt.Println("-> Proses pengurutan dengan Selection Sort selesai.")
}

func cariBerdasarkanKategori(T TabTugas, n int, keyword string) {
	var found bool = false
	fmt.Println("\nHasil Pencarian Kategori [", keyword, "]:")
	for i := 0; i < n; i++ {
		if strings.EqualFold(T[i].kategori, keyword) {
			found = true
			var statusStr string = "Belum Selesai"
			if T[i].isSelesai {
				statusStr = "Selesai"
			}
			fmt.Println("-", T[i].judul, "(DL:", T[i].deadline.tgl, "-", T[i].deadline.bln, "-", T[i].deadline.thn, ") [", statusStr, "]")
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
	target := Waktu{tgl: tgl, bln: bln, thn: thn}

	for left <= right && !found {
		mid := (left + right) / 2
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

func main() {
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
		
		pilihan = inputIntWajib("Pilih Menu (1-7): ")

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("[!] Error: Pilihan tidak valid, mohon masukkan angka antara 1-7.")
			continue
		}

		if pilihan == 1 {
			tambahTugas(&daftarTugas, &nTugas)
		} else if pilihan == 2 {
			tampilSemuaTugas(daftarTugas, nTugas)
		} else if pilihan == 3 {
			editTugas(&daftarTugas, nTugas)
		} else if pilihan == 4 {
			hapusTugas(&daftarTugas, &nTugas)
		} else if pilihan == 5 {
			fmt.Println("\n--- Sub-Menu Pencarian Data ---")
			fmt.Println("1. Cari Berdasarkan Kategori (Sequential Search)")
			fmt.Println("2. Cari Berdasarkan Tanggal Deadline (Binary Search)")
			subPil := inputIntWajib("Pilihan: ")
			
			if subPil == 1 {
				key := inputStringWajib("Masukkan Kategori yang dicari: ")
				cariBerdasarkanKategori(daftarTugas, nTugas, key)
			} else if subPil == 2 {
				fmt.Println("Catatan: Data akan otomatis diurutkan menaik sebelum eksekusi Binary Search.")
				urutBerdasarkanDeadline(&daftarTugas, nTugas, true)
				
				validSearchDL := false
				for !validSearchDL {
					inputDL := inputStringWajib("Masukkan tanggal deadline yang dicari (Tgl Bln Thn): ")
					parts := strings.Split(inputDL, " ")
					if len(parts) == 3 {
						t, errT := strconv.Atoi(parts[0])
						b, errB := strconv.Atoi(parts[1])
						th, errTh := strconv.Atoi(parts[2])
						
						if errT != nil || errB != nil || errTh != nil {
							fmt.Println("[!] Error: Format harus angka!")
						} else {
							validSearchDL = true
							resIdx := cariBerdasarkanDeadline(daftarTugas, nTugas, t, b, th)
							if resIdx != -1 {
								fmt.Println("\n-> Data ditemukan pada urutan ke-", resIdx+1, ":", daftarTugas[resIdx].judul, "[Kategori:", daftarTugas[resIdx].kategori, "]")
							} else {
								fmt.Println("-> Data tugas dengan tanggal tersebut tidak ditemukan.")
							}
						}
					} else {
						fmt.Println("[!] Error: Format tanggal salah! Gunakan spasi.")
					}
				}
			} else {
				fmt.Println("[!] Error: Pilihan sub-menu pencarian tidak valid.")
			}
		} else if pilihan == 6 {
			fmt.Println("\n--- Sub-Menu Pengurutan Data ---")
			fmt.Println("1. Urutkan Berdasarkan Tanggal Deadline (Insertion Sort)")
			fmt.Println("2. Urutkan Berdasarkan Status Kerja (Selection Sort)")
			subPil := inputIntWajib("Pilihan Kriteria: ")
			
			if subPil != 1 && subPil != 2 {
				fmt.Println("[!] Error: Pilihan kriteria tidak valid.")
				continue
			}

			fmt.Println("\nMode Pengurutan:\n1. Ascending (Menaik)\n2. Descending (Menurun)")
			mode := inputIntWajib("Pilihan Mode: ")
			
			if mode != 1 && mode != 2 {
				fmt.Println("[!] Error: Pilihan mode tidak valid.")
				continue
			}

			isAsc := true
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
		} else if pilihan == 7 {
			fmt.Println("Terima kasih telah menggunakan aplikasi! Selesai.")
		}
	}
}