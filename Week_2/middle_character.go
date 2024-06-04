package main

import (
	"fmt"
)


func main(){
	var input string
	fmt.Print("Bir Kelime Giriniz")
	fmt.Scanln(&input)

	ortaHarf:=getMiddle(input)
	fmt.Println("Metnin ortasındaki harfler:",ortaHarf)
}

func getMiddle(s string)string{
	length:=len(s)
	harfSayisiTekmiCiftmi:=length/2

	if length%2==0{
		return s[harfSayisiTekmiCiftmi-1:harfSayisiTekmiCiftmi+1]
	}else{
		return s[harfSayisiTekmiCiftmi:harfSayisiTekmiCiftmi+1]
	}
}