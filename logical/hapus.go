package logical

import (
	"fmt"
	"strings"
)

func HapusInvestasi() {
	fmt.Println("\n=== HAPUS INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	TampilkanInvestasi()

	idInput := GetInput("Masukkan ID investasi yang akan dihapus: ")
	idx := CariIndexInvestasi(idInput)

	if idx == -1 {
		fmt.Println("Investasi tidak ditemukan")
		return
	}

	// Konfirmasi hapus
	fmt.Printf("Yakin ingin menghapus investasi '%s'? (y/n): ", DataInvestasi[idx].Nama)
	konfirmasi := GetInput("")

	if strings.ToLower(konfirmasi) != "y" && strings.ToLower(konfirmasi) != "ya" {
		fmt.Println("Penghapusan dibatalkan")
		return
	}

	DataInvestasi = append(DataInvestasi[:idx], DataInvestasi[idx+1:]...)
	fmt.Println("✓ Investasi berhasil dihapus!")
}

func TampilkanInvestasi() {
	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data")
		return
	}

	fmt.Println("\nDaftar Investasi:")
	fmt.Println("--------------------------------------------------------------------------------")
	for i, inv := range DataInvestasi {
		// Tampilkan nomor urut dan ID lengkap
		fmt.Printf("%d. ID: %s\n", i+1, inv.ID)
		fmt.Printf("   Nama: %s | Jenis: %s\n", inv.Nama, inv.Jenis)
		fmt.Printf("   Dana: Rp %.2f | Nilai Kini: Rp %.2f\n", inv.Dana, inv.NilaiKini)
		fmt.Println("--------------------------------------------------------------------------------")
	}
}

func CariIndexInvestasi(idInput string) int {
	idInput = strings.TrimSpace(strings.ToLower(idInput))

	// Coba cari dengan berbagai cara
	for i, inv := range DataInvestasi {
		idLengkap := strings.ToLower(inv.ID)
		idPendek := idLengkap

		// Ambil 8 karakter pertama jika ID lebih panjang
		if len(idLengkap) > 8 {
			idPendek = idLengkap[:8]
		}

		// Cocokkan ID lengkap, ID pendek, atau mengandung input
		if idLengkap == idInput || idPendek == idInput || strings.Contains(idLengkap, idInput) {
			return i
		}
	}

	return -1
}
