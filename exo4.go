package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {

	// Lecture du fichier
	content, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	// Creation du reader CSV a partir de la variable content
	r := csv.NewReader(strings.NewReader(string(content)))

	var array []string
	cpt := 0
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if cpt != 0 {
			array = append(array, record[1])
		}
		cpt++
	}

	sort.Strings(array)
	fmt.Println(array)
}
