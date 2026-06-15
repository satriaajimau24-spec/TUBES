package logical

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const namaFile = "database_investasi.txt"

func SimpanKeFile() {
	var semuaTeks string
	for _, inv := range DataInvestasi {
		semuaTeks += fmt.Sprintf("%s,%s,%s,%.2f,%.2f\n", inv.ID, inv.Nama, inv.Jenis, inv.Dana, inv.NilaiKini)
	}
	_ = os.WriteFile(namaFile, []byte(semuaTeks), 0644)
}

func LoadDariFile() {
	isiFile, err := os.ReadFile(namaFile)
	if err != nil {
		return
	}

	baris := strings.Split(string(isiFile), "\n")
	DataInvestasi = []Investasi{}

	for _, b := range baris {
		b = strings.TrimSpace(b)
		if b == "" {
			continue
		}
		p := strings.Split(b, ",")

		if len(p) < 5 {
			continue
		}

		var inv Investasi
		inv.ID = p[0]
		inv.Nama = p[1]
		inv.Jenis = p[2]
		fmt.Sscanf(p[3], "%f", &inv.Dana)
		fmt.Sscanf(p[4], "%f", &inv.NilaiKini)

		DataInvestasi = append(DataInvestasi, inv)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetFloatInput(prompt string) float64 {
	input := GetInput(prompt)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid, menggunakan 0")
		return 0
	}
	return value
}

func FindIndexByID(id string) int {
	for i, inv := range DataInvestasi {
		if inv.ID == id || strings.HasPrefix(inv.ID, id) || inv.ID[:8] == id {
			return i
		}
	}
	return -1
}
