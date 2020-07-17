package pnml

import (
	"encoding/xml"
	"errors"
)

type IntSubtraction struct {
	XMLName xml.Name    `xml:"subtraction"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntSubtraction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intsubtraction IntSubtraction
	var ii intsubtraction
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntSubtraction: subtraction must have two subterms")
	}
	*i = IntSubtraction(ii)
	return nil
}
