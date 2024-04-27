package main

import (
	"fmt"
	"github.com/leeroyakbar/messaging-app/app"
	"github.com/leeroyakbar/messaging-app/app/maxProfit"
)

func main() {
	app.RunApp()

	//Nomor 2 -> Hitung keuntungan maksimum
	prices := []int{1, 5, 3, 8, 4, 10}
	i := 2
	profit := maxProfit.GetMaxProfit(prices, i)
	fmt.Println(profit)
}
