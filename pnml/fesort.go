package pnml

import "encoding/xml"

type FESort struct {
	XMLName   xml.Name     `xml:"finiteenumeration"`
	Constants []FEConstant `xml:"feconstant"` // optional
}
