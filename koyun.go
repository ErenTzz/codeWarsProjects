package main

import (
	"fmt"
)

func main() {
	sheepNumber := getUserInput()
	for i := 1; i <= sheepNumber; i++ {
		fmt.Printf("%d koyun... ", i)
	}
	fmt.Println()
}

func getUserInput() int {
	var input int
	for {
		fmt.Print("Kaç koyun saymak istediğinizi giriniz: ")
		_, err := fmt.Scan(&input)// inputta herhangi bir hata olup olmadığını kontrol eder
		if err != nil || input <= 0 {
			fmt.Println("Geçersiz giriş. Lütfen pozitif bir tam sayı giriniz.")

			var discard string //geçersiz girişi tutar
			fmt.Scanln(&discard)// ve discard eder
		} else {
			fmt.Println("İstediğin koyun miktarı:", input)
			break
		}
	}
	return input
}
