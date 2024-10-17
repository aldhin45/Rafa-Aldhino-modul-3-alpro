package main

import (
	"fmt"
)

func combined(a, b, c int) (fogoh int, gohof int, hofog int) {
	fogoh = (a - 1) * (a - 1)
	gohof = (b - 1) - 2
	hofog = ((c - 1) * (c - 1)) + 1
	return
}

func main() {
	var a, b, c int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Menghitung hasil dengan fungsi combined
	fogoh, gohof, hofog := combined(a, b, c)

	// Menampilkan hasil
	fmt.Println("Hasil fogoh:", fogoh)
	fmt.Println("Hasil gohof:", gohof)
	fmt.Println("Hasil hofog:", hofog)
}
