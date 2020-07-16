package pnml

import "encoding/xml"

type CyclicEnumSort struct {
	XMLName   xml.Name     `xml:"cyclicenumeration"`
	Constants []FEConstant `xml:"feconstant"` // optional
}
