package main

import (
	"fmt"
)

func main() {
	anapara := 0.0
	aylikOdeme := 200.0
	aySayisi := 144

	fmt.Println(anapara, aylikOdeme)

	aylikoran := aylikOran(3.0, 0.18)
	aylikkazanc := aylikKazanc(aylikOdeme, aylikoran)

	fmt.Println("aylikkazanc", aylikkazanc, "aylikoran", aylikoran)
	for i := 0; i < aySayisi; i++ {

		if i == 0 {
			anapara += aylikOdeme
		} else {
			anapara += aylikOdeme + aylikKazanc(anapara, aylikoran)
		}

		// Mod(aylikKazanc(aylikOdeme, aylikOran(3.0, 0.18)), 1.0)
	}
	fmt.Println("Odenen :", int(aylikOdeme)*aySayisi, " - toplam :",
		anapara, " - fark :", int(anapara)-(int(aylikOdeme)*aySayisi),
		" - %  kazanc :", 100-float64(100)*(float64(aylikOdeme)*float64(aySayisi))/float64(anapara))

	fmt.Println("Aegonda Ödenen :", float64(aySayisi)*aylikOdeme, " - Reel Ödenen :",
		float64(aySayisi)*aylikOdeme*0.7)
}

func aylikOran(yillikOran, stopaj float64) float64 {
	return (yillikOran / 12) * (1.0 - stopaj) / 100.0
}

func aylikKazanc(para, aylıkoran float64) float64 {
	return para * aylıkoran
}
