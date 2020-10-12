package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("IntMultiplication: mult must have two subterms")
		}
		log.Print("Pinimili: mult element with ", len(ii.Terms), " subterm elements (should be 2)")
	}
	*i = IntMultiplication(ii)
	return nil
}
