package logical

import (
	"fmt"
)

func UbahInvestasi() {
	fmt.Println("\n=== UBAH INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	TampilkanInvestasi()

	id := GetInput("Masukkan ID investasi yang akan diubah: ")
	idx := CariIndexInvestasi(id)

	if idx == -1 {
		fmt.Println("Investasi tidak ditemukan")
		return
	}

	fmt.Printf("\nMengubah data: %s\n", DataInvestasi[idx].Nama)
	fmt.Print("Nama Aset baru: ")
	fmt.Scanln(&DataInvestasi[idx].Nama)

	fmt.Print("Jenis Aset baru: ")
	fmt.Scanln(&DataInvestasi[idx].Jenis)

	fmt.Print("Jumlah Dana baru: ")
	fmt.Scanln(&DataInvestasi[idx].Dana)

	fmt.Print("Nilai Terkini baru: ")
	fmt.Scanln(&DataInvestasi[idx].NilaiKini)

	fmt.Println("✓ Investasi berhasil diubah!")
}
