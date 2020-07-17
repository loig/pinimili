package pnml

import (
	"encoding/xml"
	"errors"
)

type StringLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLessThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlessthanorequal StringLessThanOrEqual
	var ss stringlessthanorequal
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringLessThanOrEqual: leqs must have two subterms")
	}
	*s = StringLessThanOrEqual(ss)
	return nil
}
