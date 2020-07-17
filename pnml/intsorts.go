package pnml

import "encoding/xml"

type IntNatural struct {
	XMLName xml.Name `xml:"natural"`
}

type IntPositive struct {
	XMLName xml.Name `xml:"positive"`
}

type IntInteger struct {
	XMLName xml.Name `xml:"integer"`
}
