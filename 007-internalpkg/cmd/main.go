package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type myStruct struct {
	a int
	b string
}

func main() {
	fmt.Println("Basliyooo ....")
	// zaman.GunlerdenNe()

	var bir myStruct
	iki := myStruct{1, "Iki"}
	uc := new(myStruct)

	log.Info(bir, iki, *uc)
}
