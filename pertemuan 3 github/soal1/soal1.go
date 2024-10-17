package main

import "fmt"

func main() {
	var n, r int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)

	fmt.Println("Hasil Permutasi:", permutasi(n, r))
}

func mencariFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := n; i >= 1; i-- {
		*hasil = *hasil * i
	}
}

func permutasi(n, r int) int {
	var pembilang, penyebut int
	mencariFaktorial(n, &pembilang)
	mencariFaktorial(n-r, &penyebut)
	return pembilang / penyebut
}

func faktorial(n int) int {
	hasil := 1
	for i := n; i >= 1; i-- {
		hasil *= i
	}
	return hasil
}

func combinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
