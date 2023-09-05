package main

import "fmt"

// function parameter

func printAngka(angka1 int, angka2 int) {
	fmt.Println("ini baris", angka1)
	fmt.Println("ini baris", angka2)
}

// function return value

func introduction(name string) string {
	return "Hello my name is " + name

}

func main() {

	fmt.Println("Initial awal")
	fmt.Println("")

	// function parameter
	printAngka(1, 2)
	fmt.Println("")

	// function return value
	// --panggil langsung
	fmt.Println(introduction("ralif"))
	// --panggil menggunakan variabel
	result := introduction("ralif")
	fmt.Println(result)
}
