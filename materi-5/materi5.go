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

func introduction2(firstName string, lastName string) (string, string) {

	introFirstName := "Hello my first name is " + firstName
	introLastName := "Hello my last name is " + lastName

	return introFirstName, introLastName

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

	// function sebagai value
	secondResult := introduction
	fmt.Println(secondResult("ralif"))

	// function return multiple value
	fmt.Println("")

	firstName, lastname := introduction2("Reza", "Alif")
	fmt.Println(firstName, "\n", lastname)

	// jika tidak ingin menggunakan value bisa menggunakan tanda _
	fmt.Println("")
	firstName2, _ := introduction2("Reza", "Alif")
	fmt.Println(firstName2)
}
