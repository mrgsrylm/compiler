package main

import (
	"fmt"
)

func main() {
	characters := `САШАРВО`

	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %v \n", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			for i, value := range characters {
				fmt.Printf("character %#U starts at byte position %d \n", value, i)
			}
			continue
		}else{
			fmt.Printf("Nilai j = %v \n", j)
		}
	}

}
