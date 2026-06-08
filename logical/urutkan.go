package logical

import (
	"fmt"
	"sort"
)

func UrutkanInvestasi() {
	fmt.Println("\n=== URUTKAN INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	fmt.Println("1. Berdasarkan Nama (Built-in)")
	fmt.Println("2. Berdasarkan Jenis (Built-in)")
	fmt.Println("3. Berdasarkan Dana (Built-in)")
	fmt.Println("4. Berdasarkan Nilai Kini (Built-in)")
	fmt.Println("5. Selection Sort (Nilai Kini)")
	fmt.Println("6. Insertion Sort (Nilai Kini)")

	pilihan := GetInput("Pilih pengurutan (1-6): ")

	switch pilihan {

	// ===== BUILT-IN SORT =====
	case "1":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Nama < DataInvestasi[j].Nama
		})

	case "2":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Jenis < DataInvestasi[j].Jenis
		})

	case "3":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Dana < DataInvestasi[j].Dana
		})

	case "4":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].NilaiKini < DataInvestasi[j].NilaiKini
		})

	// ===== SELECTION SORT =====
	case "5":
		selectionSort(DataInvestasi)

	// ===== INSERTION SORT =====
	case "6":
		insertionSort(DataInvestasi)

	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Println("✓ Data berhasil diurutkan")
	TampilkanInvestasi()
}

// =====================
// SELECTION SORT
// =====================
func selectionSort(data []Investasi) {
	n := len(data)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].NilaiKini < data[minIdx].NilaiKini {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

// =====================
// INSERTION SORT
// =====================
func insertionSort(data []Investasi) {
	n := len(data)

	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].NilaiKini > key.NilaiKini {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
