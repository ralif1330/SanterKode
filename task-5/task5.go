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

func introduce() (nama string, gender string, job string, umur int, ex string) {
	nama = "Pak John"
		gender = "adalah seorang"
			job = "penulis yang berusia"
				umur = 30
					ex = "tahun"

	return
}

func introduce2() (nama string, gender string, job string, umur int, ex string) {
	nama = "Bu Sarah"
		gender = "adalah seorang"
			job = "model yang berusia"
				umur = 28
					ex = "tahun"

	return
}

// soal 3






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

	john, gender, work, age, ex := introduce()
		fmt.Println(john, gender, work, age, ex) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah, gender, work, age, ex := introduce2()
		fmt.Println(sarah, gender, work, age, ex) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3

}