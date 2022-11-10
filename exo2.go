package main

import (
	"fmt"
)


func main() {

	var transactions uint
	var time uint
	var format [3]string
	format[0] = "heures"
	format[1] = "minutes"
	format[2] = "secondes"

	fmt.Println("Veuillez entrer le nombre de transactions")
	fmt.Scanln(&transactions)

	choice := 0
	fmt.Println("Veuillez choisir le format souhaité")

	for ; choice<1 || choice >3 ; {

		fmt.Println("1 - Heures")
		fmt.Println("2 - Minutes")
		fmt.Println("3 - Secondes")
		fmt.Scanln(&choice)

		if choice<1 || choice>3 {
			fmt.Println("Veuillez choisir un chiffre entre 1 et 3 pour le format")
		}

	}


	fmt.Println("Veuillez entrer le temps au format en", format[choice-1])
	fmt.Scanln(&time)

	result := float32(transactions) / float32(time)
	fmt.Println("Il y a eu", result, "transactions par", format[choice-1])

	/*switch choice {
	case 1:
		result := transactions / 60
		fmt.Println("Il y a eu", result, "transactions par heure")
	case 2:
		result := transactions / 60
		fmt.Println("Il y a eu", result, "transactions par minute")
	default:
		result := transactions / 60
		fmt.Println("Il y a eu", result, "transactions par seconde")
	}*/



	/*split := strings.Fields(time)

	fmt.Println(split)
	fmt.Println(time)*/






}