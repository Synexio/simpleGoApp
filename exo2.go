package main

import (
	"fmt"
	"time"
)

func main() {
	var age uint16
	var date time.Time
	var birthYear uint16

	fmt.Println("Bonjour, veuillez entrer votre age en annee")
	fmt.Scanf("%d\n", &age)

	date = time.Now()
	birthYear = uint16(date.Year()) - age
	fmt.Println("Nous sommes en", date.Year())
	fmt.Println("Vous etes née en", birthYear, "ou bien en", birthYear-1)
}
