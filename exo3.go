package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	var addr string

	fmt.Println("Veuillez rentrer une URL")
	fmt.Scanf("%s\n", &addr)

	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme :", u.Scheme)
	fmt.Println("Host :", u.Host)
	fmt.Println("Path :", u.Path)
	fmt.Println("Fragment :", u.Fragment)
}
