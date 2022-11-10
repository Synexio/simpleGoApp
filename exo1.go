package main

import "fmt"


type Ville struct{
	name string
	zipCode uint16
	numberPeople uint32
	size uint32
}

func main() {
	fmt.Println("Bonjour, veuillez entrer 3 villes")

	m := make(map[int]Ville)
	var name string
	var zipCode uint16
	var numberPeople uint32
	var size uint32

	for i:=0; i<3; i++ {
		fmt.Println("Ville numéro ", i)

		fmt.Println("Nom :")
		fmt.Scanf("%s\n",&name)

		fmt.Println("Code postal :")
		fmt.Scanf("%d\n",&zipCode)

		fmt.Println("Nombre d'habitants :")
		fmt.Scanf("%d\n",&numberPeople)

		fmt.Println("Superficie :")
		fmt.Scanf("%d\n",&size)

		m[i] = Ville{name, zipCode, numberPeople, size}

		defer fmt.Println(m[i])
	}



}
