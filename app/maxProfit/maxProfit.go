package maxProfit

import (
	"math"
)

// function untuk memperoleh keuntungan maksimal dengan i transaksi
func GetMaxProfit(prices []int, k int) int {
	n := len(prices)
	// membuat slice untuk menyimpan keuntungan maksimum setelah membeli
	profitAfterBuying := make([]int, k+1)
	// membuat slice untuk menyimpan keuntungan maksimum setelah membeli
	profitAfterSelling := make([]int, k+1)

	// Inisialisasi profitAfterBuying dengan nilai minimum
	for j := range profitAfterBuying {
		profitAfterBuying[j] = math.MinInt32
	}

	// Iterasi pada setiap hari
	for i := 0; i < n; i++ {
		currentPrice := prices[i]
		for j := k; j >= 1; j-- {
			// Hitung keuntungan maksimum setelah membeli
			profitAfterBuying[j] = max(profitAfterBuying[j], profitAfterSelling[j-1]-currentPrice)
			// Hitung keuntungan maksimum setelah menjual
			profitAfterSelling[j] = max(profitAfterSelling[j], profitAfterBuying[j]+currentPrice)
		}
	}

	// return profit maksimum
	return profitAfterSelling[k]
}

// function max untuk mencari nilai terbesar
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
