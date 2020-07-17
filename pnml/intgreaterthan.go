package pnml

import (
	"encoding/xml"
	"errors"
)

type IntGreaterThan struct {
	XMLName xml.Name    `xml:"gt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intgreaterthan IntGreaterThan
	var ii intgreaterthan
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntGreaterThan: gt must have two subterms")
	}
	*i = IntGreaterThan(ii)
	return nil
}
