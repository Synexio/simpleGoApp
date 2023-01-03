package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Creation fichier
	message := []byte("Hello world!")
	// err := os.WriteFile("hello.txt", message, 0644)
	err := ioutil.WriteFile("hello.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	// Lecture fichier
	content, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	// Affichage
	fmt.Println(string(content))
}
