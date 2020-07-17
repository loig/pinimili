package pnml

import (
	"encoding/xml"
	"errors"
)

type IntLessThan struct {
	XMLName xml.Name    `xml:"lt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intlessthan IntLessThan
	var ii intlessthan
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntLessThan: lt must have two subterms")
	}
	*i = IntLessThan(ii)
	return nil
}
