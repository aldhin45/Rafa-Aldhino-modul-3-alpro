package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func isInsideCircle(cx, cy, r, px, py float64) bool {
	return distance(cx, cy, px, py) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, px, py float64

	fmt.Print("Masukkan koordinat x dan y pusat lingkaran 1: ")
	fmt.Scan(&cx1, &cy1)
	fmt.Print("Masukkan jari-jari lingkaran 1: ")
	fmt.Scan(&r1)

	fmt.Print("Masukkan koordinat x dan y pusat lingkaran 2: ")
	fmt.Scan(&cx2, &cy2)
	fmt.Print("Masukkan jari-jari lingkaran 2: ")
	fmt.Scan(&r2)

	fmt.Print("Masukkan koordinat x dan y titik: ")
	fmt.Scan(&px, &py)

	inCircle1 := isInsideCircle(cx1, cy1, r1, px, py)
	inCircle2 := isInsideCircle(cx2, cy2, r2, px, py)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
