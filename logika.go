package main

import "fmt"

func main() {
	var nilaiSiswa = 75

	var statusNilai = nilaiSiswa > 60

	// penggunaan operator
	var siswaDapatUjian = !statusNilai

	// output
	fmt.Print(siswaDapatUjian)
}
