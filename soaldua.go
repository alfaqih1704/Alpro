package main

import "fmt"

func main() {
	var nama, nim, kelas string

	// Membaca masukan: nama, nim, dan kelas
	fmt.Print("masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("masukan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("masukan kelas: ")
	fmt.Scanln(&kelas)


	// Menampilkan resume singkat mahasiswa
	fmt.Println()
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}