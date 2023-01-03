package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Urls struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"urlset"`
	Loc     string   `xml:"loc"`
	Lastmod string   `xml:"lastmod"`
}

func main() {
	// Make a GET request to the sitemap
	response, err := http.Get("https://www.esgi.fr/page-sitemap.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the sitemap into the appropriate struct
	var sitemap Urls
	err = xml.Unmarshal(body, &sitemap)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the locations
	for _, x := range sitemap.Urls {
		fmt.Println(x.Loc + "modifié à" + x.Lastmod)
	}

}
