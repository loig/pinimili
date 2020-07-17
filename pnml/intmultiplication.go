package pnml

import (
	"encoding/xml"
	"errors"
)

type IntMultiplication struct {
	XMLName xml.Name    `xml:"mult"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntMultiplication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intmultiplication IntMultiplication
	var ii intmultiplication
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntMultiplication: mult must have two subterms")
	}
	*i = IntMultiplication(ii)
	return nil
}
