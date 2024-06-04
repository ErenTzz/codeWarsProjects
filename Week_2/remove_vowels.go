package main

import (
	"fmt"
	"strings"
)


var vowels = "aeiouAEIOU"


func removeVowels(s string) string {
	var result strings.Builder
	for _, char := range s {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {

	input := "Selamlar Arkadaşlar!"
	output := removeVowels(input)
	fmt.Println("Girdi:", input)
	fmt.Println("Çıktı:", output)
}
