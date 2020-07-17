package pnml

import (
	"encoding/xml"
	"errors"
)

type IntDivision struct {
	XMLName xml.Name    `xml:"div"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntDivision) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intdivision IntDivision
	var ii intdivision
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntDivision: div must have two subterms")
	}
	*i = IntDivision(ii)
	return nil
}
