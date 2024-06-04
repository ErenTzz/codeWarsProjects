package main

import (
	"fmt"
	"strings"
	"unicode"
)


func ilkHarfBuyuk(buyuk string) string {
	kelimeler := strings.Fields(buyuk)
	for i, kelime := range kelimeler {
		kelimeler[i] = string(unicode.ToUpper(rune(kelime[0]))) + kelime[1:]
	}
	return strings.Join(kelimeler, " ")
}

func main() {
	input := "merhaba beyler nasılsınız iyi misiniz ?"
	output := ilkHarfBuyuk(input)
	fmt.Println(output)
}