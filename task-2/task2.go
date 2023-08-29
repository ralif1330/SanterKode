package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}

func soal1() {

	kataPertama := "Bootcamp "
	kataKedua := "Digital "
	kataKetiga := "Skill "
	kataKeEmpat := "SanterKode "
	kataKelima := "Golang"

	fmt.Print(kataPertama, kataKedua, kataKetiga, kataKeEmpat, kataKelima)

}

// Menggunakan Package Strings.Replace

func soal2() {

	kata := "Halo Dunia"
	find := "Dunia"
	ReplaceWith := "Golang"

	newTextSoal2 := strings.Replace(kata, find, ReplaceWith, 1)
	fmt.Println("")
	fmt.Println(newTextSoal2)

}

func soal3() {

	var kataPertama = "saya "
	var kataKedua = "senang "
	var kataKetiga = "belajar "
	var kataKeempat = "golang"

	// Kata Pertamaa

	fmt.Print(kataPertama)

	// kata Kedua

	updateKedua := strings.Title(kataKedua)
	fmt.Print(updateKedua)

	// kata ketiga

	find3 := "r"
	replace3With := "R"
	newKata3 := strings.Replace(kataKetiga, find3, replace3With, 1)
	fmt.Print(newKata3)

	// kata keEmpat

	upperCase := strings.ToUpper(kataKeempat)
	fmt.Println(upperCase)

}

func soal4() {

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var cv1, _ = strconv.Atoi(angkaPertama)
	var cv2, _ = strconv.Atoi(angkaKedua)
	var cv3, _ = strconv.Atoi(angkaKetiga)
	var cv4, _ = strconv.Atoi(angkaKeempat)

	total := cv1 + cv2 + cv3 + cv4
	fmt.Print(total)

}

func soal5() {

	kalimat := "halo halo bandung"
	angka := 2021

	find := "halo halo "
	ReplaceWith := "Hi Hi "

	newTextSoal5 := strings.Replace(kalimat, find, ReplaceWith, 1)
	fmt.Println("")
	fmt.Println(newTextSoal5, "-", angka)

}
