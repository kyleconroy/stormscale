package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Mechanic struct {
	Effect   string   `json:"effect"`
	RulesText string   `json:"rule"`
	Blocks   []string `json:"blocks"`
	CardURL  string   `json:"example"`
	Rating   int      `json:"rating"`
}

type Mechanics []Mechanic

type Payload struct {
	Scale []string
	Items Mechanics
}

func main() {
	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
		log.Fatal(err)
	}

	var mechanics Mechanics
	var scale []string

	blob, err := ioutil.ReadFile("stormscale.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(blob, &scale)

	if err != nil {
		log.Fatal(err)
	}

	blob, err = ioutil.ReadFile("mechanics.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(blob, &mechanics)

	writer, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}

	payload := Payload{Items: mechanics, Scale: scale}

	err = tmpl.Execute(writer, payload)

	if err != nil {
		log.Fatal(err)
	}

}
