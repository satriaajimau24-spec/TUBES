package logical

import (
	"fmt"

	"github.com/google/uuid"
)

func TambahInvestasi() {
	var inv Investasi

	fmt.Println("\n=== TAMBAH INVESTASI ===")
	inv.ID = uuid.New().String()

	fmt.Print("Nama Aset: ")
	fmt.Scanln(&inv.Nama)

	fmt.Print("Jenis Aset (Saham/Obligasi/ReksaDana): ")
	fmt.Scanln(&inv.Jenis)

	fmt.Print("Jumlah Dana: ")
	fmt.Scanln(&inv.Dana)

	fmt.Print("Nilai Terkini: ")
	fmt.Scanln(&inv.NilaiKini)

	DataInvestasi = append(DataInvestasi, inv)
	fmt.Println("✓ Investasi berhasil ditambahkan!")
}
