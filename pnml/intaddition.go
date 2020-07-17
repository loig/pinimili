package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("IntAddition: addition must have two subterms")
	}
	*i = IntAddition(ii)
	return nil
}
