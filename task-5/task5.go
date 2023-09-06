package main

import "fmt"

// soal 1

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return (panjang + lebar) * 2
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// soal 2

func introduce() (nama string, sex string, job string, umur int) {
	nama = "John"
	sex = "laki-laki"
	job = "penulis"
	umur = 30

	return
}

func introduce2() (nama string, sex string, job string, umur int) {
	nama = "Sarah"
	sex = "perempuan"
	job = "model"
	umur = 28

	return
}

func main() {

	// soal 1

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("hasil luas persegi panjang : ", luas)
	fmt.Println("hasil keliling persegi panjang : ", keliling)
	fmt.Println("hasil volume balok : ", volume)
	fmt.Println("")

	// soal 2

	john, gender, work, age := introduce()
	fmt.Println(john, gender, work, age) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah, gender, work, age := introduce2()
	fmt.Println(sarah, gender, work, age) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

}
