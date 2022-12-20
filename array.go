package main

import "fmt"

func main() {

	/**
	* Deklarasi array variabel dengan panjang array 3 dengan isi nilai bertipe data string
	 */
	// var animals [3]string

	// //pengisian array
	// animals[0] = "Kambing"
	// animals[1] = "Sapi"
	// animals[2] = "ayam"
	// // output
	// // fmt.Println(animals)

	// animals[1] = "n"

	// fmt.Println(animals)

	var animals = [3]string{
		"Kambing",
		"Sapi",
		"Ayam",
	}

	fmt.Println(animals)

}
