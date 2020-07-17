package pnml

import (
	"encoding/xml"
	"errors"
)

type IntGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntGreaterThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intgreaterthanorequal IntGreaterThanOrEqual
	var ii intgreaterthanorequal
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntGreaterThanOrEqual: geq must have two subterms")
	}
	*i = IntGreaterThanOrEqual(ii)
	return nil
}
