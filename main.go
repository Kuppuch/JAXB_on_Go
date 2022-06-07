package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
}

func main() {
	book := Book{
		ID:   1,
		Name: "кНига",
	}
	byteBook, err := xml.MarshalIndent(book, " ", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(byteBook))
}
