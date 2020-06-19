package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	pnmlFile, err := os.Open("test.pnml")
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(pnmlFile)
	if err != nil {
		panic(err)
	}

	var pnml Pnml
	err = xml.Unmarshal(byteValue, &pnml)
	if err != nil {
		panic(err)
	}

	fmt.Println("This model has", len(pnml.Nets), "nets")
	for i, net := range pnml.Nets {
		fmt.Println("Net", i, "has", len(net.Pages), "pages and is called", net.Name.Text)
		fmt.Println("Its type is", net.Type)
		for j, page := range net.Pages {
			fmt.Println("Page", j, "\t", len(page.Pages), "pages")
			fmt.Println("\t", len(page.Places), "places")
			fmt.Println("\t", len(page.Transitions), "transitions")
			fmt.Println("\t", len(page.Arcs), "arcs")
		}
	}

	byteValue2, err := xml.Marshal(pnml)
	fmt.Println(string(byteValue2))
}
