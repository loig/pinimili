package pnml

import (
	"encoding/xml"
	"errors"
)

type StringGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringGreaterThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringgreaterthanorequal StringGreaterThanOrEqual
	var ss stringgreaterthanorequal
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringGreaterThanOrEqual: geqs must have two subterms")
	}
	*s = StringGreaterThanOrEqual(ss)
	return nil
}
