package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	// Membaca masukan jari-jari lingkaran
	fmt.Scan(&r)

	// Menhitung luas lingkaran (L = pi * r^2)
	luas := math.Pi * r * r

	// Menampilkan keluaran dengan 1 angka di belakang koma
	fmt.Printf("%.1f\n", luas)
}