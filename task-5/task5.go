package main

import "fmt"

func main() {
	fmt.Println("tes 1 bcxe ")

	printHello()

	fmt.Println(introduction("Reza"))
}

// function parameter pertamam

func printHello() {
	fmt.Println("Hallo ini baris pertama nyoba function")
}

// function return value

func introduction(name string) string {
	return "Hello my name is " + name
}
