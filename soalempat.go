package main

import "fmt"

func main() {
	var f float64

	// Membaca masukan suhu dalam Fahrenheit
	fmt.Scan(&f)

	// Menghitung konversi ke Celsius
	c := (f - 32) * 5 / 9

	// Menampilkan hasil suhu dalam Celsius
	fmt.Printf("%.0f\n", c)
}