package main

import "fmt"

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()

}

func soal1() {

	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d I Love Coding\n", i)
		} else if i%2 != 0 {
			fmt.Printf("%d Santai\n", i)
		} else {
			fmt.Printf("%d Berkualitas\n", i)
		}
	}
}

func soal2() {
	fmt.Println("")
	rows := 7

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func soal3() {
	fmt.Println("")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	aKalimat := kalimat[2:7]
	fmt.Println(aKalimat)

}

func soal4() {
	fmt.Println("")
	var sayuran = make([]string, 0)

	sayuran = append(sayuran, "bayam", "buncis", "kangkung", "kubis", "seledri", "tauge", "timun")

	fmt.Println(sayuran)

	for i, sayur := range sayuran {
		fmt.Printf(" %d. %s \n", i+1, sayur)
	}

}

func soal5() {
	fmt.Println("")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	i := 1
	for key, value := range satuan {
		fmt.Printf("%s: %d\n", key, value)
		i++
	}

	volumes := []int{}

	panjang := 7
	lebar := 4
	tinggi := 6

	volume := panjang * lebar * tinggi
	volumes = append(volumes, volume)

	// Menampilkan hasil perhitungan
	fmt.Println("Volume balok:", volumes)
}
