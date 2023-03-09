package main

import (
	"fmt"
	"gofund/chap4/helpers"
	"os"
	"strconv"
)

func main() {
	data_student := helpers.Init()
	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			absen, _ := strconv.Atoi(os.Args[i])

			if absen <= 0 {
				fmt.Printf("\nMasukan Urutan %v Tidak Valid!\n", absen)
				continue
			}

			if absen >= len(data_student) {
				fmt.Printf("\nBiodata Urutan %v Tidak Ditemukan!\n", absen)
				continue
			}

			helpers.Print(data_student, absen)
		}
	} else {
		fmt.Println("\nMasukan Urutan Biodata Setelah biodata.go!")
	}
}
