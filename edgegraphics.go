package main

import "encoding/xml"

type EdgeGraphics struct {
	XMLName   xml.Name   `xml:"graphics"`
	Positions []Position `xml:"position"` // optional
	Line      *Line      `xml:"line"`     // optional
}
