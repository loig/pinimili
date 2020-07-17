package pnml

import (
	"encoding/xml"
	"errors"
)

type IntLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntLessThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intlessthanorequal IntLessThanOrEqual
	var ii intlessthanorequal
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		return errors.New("IntLessThanOrEqual: leq must have two subterms")
	}
	*i = IntLessThanOrEqual(ii)
	return nil
}
