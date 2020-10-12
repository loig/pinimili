package pnml

import (
	"encoding/xml"
	"errors"
	"log"
)

type IntAddition struct {
	XMLName xml.Name    `xml:"addition"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntAddition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intaddition IntAddition
	var ii intaddition
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		if panicIfNotPnmlCompliant {
			return errors.New("IntAddition: addition must have two subterms")
		}
		log.Print("Pinimili: addition element with ", len(ii.Terms), " subterm elements (should be 2)")
	}
	*i = IntAddition(ii)
	return nil
}
