package main

import (
	"fmt"
	"sort"
	"strings"
)


func main() {

	var size uint8
	var cpt uint8
	var list []string
	fmt.Println("Veuillez choisir le nombre d'entrées dans la liste")
	fmt.Scanln(&size)

	for cpt=0; cpt<size; cpt++ {
		var input string
		fmt.Println("Veuillez rentrer le mot numero", cpt+1)
		fmt.Scanln(&input)
		list = append(list, strings.ToLower(input))
	}

	sort.Strings(list)

	fmt.Println(list)

}