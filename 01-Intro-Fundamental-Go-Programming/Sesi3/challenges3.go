package main

import (
	"fmt"

	"gofund/chap3/helpers"
)

func main() {
	statement := helpers.ParseStrings("selamat malam")
	fmt.Println(statement)
}

