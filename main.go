package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/loig/pinimili/pnml"
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

	var pnml pnml.Pnml
	err = xml.Unmarshal(byteValue, &pnml)
	if err != nil {
		panic(err)
	}

	fmt.Println("Unmarshaling ok")

	/*
		fmt.Println("This model has", len(pnml.Nets), "nets")
		for i, net := range pnml.Nets {
			fmt.Println("Net", i, "has", len(net.Pages), "pages and", len(net.HLDeclarations), "high level declarations")
			fmt.Println("Its type is", *net.Type)
			for j, decl := range net.HLDeclarations {
				fmt.Println("High level declaration", j)
				fmt.Println("\t", len(decl.SortDeclarations), "sorts")
				fmt.Println("\t", len(decl.VariableDeclarations), "variables")
				fmt.Println("\t", len(decl.OperatorDeclarations), "operators")
			}
			for j, page := range net.Pages {
				fmt.Println("Page", j)
				fmt.Println("\t", len(page.Pages), "pages")
				fmt.Println("\t", len(page.Places), "places")
				fmt.Println("\t", len(page.Transitions), "transitions")
				fmt.Println("\t", len(page.Arcs), "arcs")
			}
		}
	*/

	/*
		net := ptnet.NewFromPnml(pnml)
		fmt.Println("Extracted net has:")
		fmt.Println("\t", len(net[0].Places), "places")
		fmt.Println("\t", len(net[0].Transitions), "transitions")
		fmt.Println("\t", len(net[0].Arcs), "arcs")
	*/

	/*
		byteValue2, err := xml.MarshalIndent(pnml, "", "\t")
		fmt.Println(string(byteValue2))
	*/
}
