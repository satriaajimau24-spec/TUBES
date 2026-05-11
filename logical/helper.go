package logical

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		// Cocokkan dengan ID lengkap atau 8 karakter pertama
		if inv.ID == id || strings.HasPrefix(inv.ID, id) || inv.ID[:8] == id {
			return i
		}
	}
	return -1
}
