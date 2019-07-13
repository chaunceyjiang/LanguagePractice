package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("default.xml")
	if err != nil {
		log.Println(err)
	}
	decode := xml.NewDecoder(file)
	for t, err := decode.Token(); err != nil; t, err = decode.Token() {
		switch token := t.(type) {
		case xml.StartElement:
		case xml.EndElement:

		}
	}
	fmt.Println()
}
