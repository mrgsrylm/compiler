package helpers

import "fmt"

func ParseStrings(param string) map[string]int {
	var characters []string

	for _, char := range param {
		characters = append(characters, string(char))
	}
	
	dictionary := make(map[string]int)
	for _, char := range characters {
		fmt.Println(char)
		dictionary[char]++
	}

	return dictionary
}