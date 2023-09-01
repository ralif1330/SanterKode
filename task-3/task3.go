package main

import (
	"fmt"
	"strconv"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()

}

func soal1() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// ubah string ke int

	var ppp, _ = strconv.Atoi(panjangPersegiPanjang)
	var lpp, _ = strconv.Atoi(lebarPersegiPanjang)
	var als, _ = strconv.Atoi(alasSegitiga)
	var ts, _ = strconv.Atoi(tinggiSegitiga)

	fmt.Print(
		"Panjang Persegi Panjang : ", ppp,
		"\nLebar Persegi Panjang : ", lpp,
		"\nAlas Segitiga : ", als,
		"\nTinggi Segitiga : ", ts,
	)
	fmt.Println("")

	var luasPersegiPanjang int = ppp * lpp
	var kelilingPersegiPanjang int = 2 * (ppp + lpp)
	var luasSegitiga int = (als * ts) / 2

	fmt.Printf("Rumus Luas Persegi Panjang = Panjang x Lebar  = %d\n", luasPersegiPanjang)
	fmt.Printf("Rumus Keliling Persegi Panjang = (Panjang + Lebar) x 2 = %d\n", kelilingPersegiPanjang)
	fmt.Printf("Rumus Luas Segitiga = (alas x tinggi ) / 2 = %d\n", luasSegitiga)

}

func soal2() {

	nilaiJohn := 80
	nilaiDoe := 50

	var indeksJohn string
	var indeksDoe string

	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}

	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}

	fmt.Println("")
	fmt.Printf("Nilai John: %d\nIndeks: %s\n", nilaiJohn, indeksJohn)
	fmt.Printf("Nilai Doe: %d\nIndeks: %s\n", nilaiDoe, indeksDoe)
}

func soal3() {

	var tanggal = 13
	var bulan = 03
	var tahun = 2000
	namaBulan := ""

	switch bulan {
	case 01:
		namaBulan = "Januari"
	case 02:
		namaBulan = "Februari"
	case 03:
		namaBulan = "Maret"
	default:
		namaBulan = "Bulan Tidak Ada"
	}
	fmt.Printf("\n%d %s %d \n", tanggal, namaBulan, tahun)
	fmt.Println("")
}

func soal4() {

	tahunLahir := 2000

	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		fmt.Println("Termasuk dalam Generasi Baby Boomer")
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		fmt.Println("Termasuk dalam Generasi X")
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		fmt.Println("Termasuk dalam Generasi Y (Millenials)")
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		fmt.Println("Termasuk dalam Generasi Z")
	} else {
		fmt.Println("Tahun kelahiran tidak masuk dalam kategori yang ada")
	}
}
