package helpers

import (
	"fmt"
)

func Print(data_student []student, absen int) {
	fmt.Println("Nama\t\t\t\t: ", data_student[absen].nama)
	fmt.Println("Alamat\t\t\t\t: ", data_student[absen].alamat)
	fmt.Println("Pekerjaan\t\t\t: ", data_student[absen].pekerjaan)
	fmt.Println("Alasan memilih kelas Golang\t: ", data_student[absen].alasan)
}