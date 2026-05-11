package logical

import (
	"fmt"
)

func LaporanPortofolio() {
	fmt.Println("\n=== LAPORAN PORTOFOLIO ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	saham := 0.0
	obligasi := 0.0
	reksadana := 0.0

	for _, inv := range DataInvestasi {
		switch inv.Jenis {
		case "Saham":
			saham += inv.NilaiKini
		case "Obligasi":
			obligasi += inv.NilaiKini
		case "ReksaDana":
			reksadana += inv.NilaiKini
		}
	}

	total := saham + obligasi + reksadana

	fmt.Printf("\nTotal Portofolio: %.2f\n", total)
	fmt.Printf("Saham: %.2f (%.1f%%)\n", saham, (saham/total)*100)
	fmt.Printf("Obligasi: %.2f (%.1f%%)\n", obligasi, (obligasi/total)*100)
	fmt.Printf("ReksaDana: %.2f (%.1f%%)\n", reksadana, (reksadana/total)*100)
}
