package pnml

import (
	"encoding/xml"
	"errors"
)

type StringLessThan struct {
	XMLName xml.Name    `xml:"lts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlessthan StringLessThan
	var ss stringlessthan
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringLessThan: lts must have two subterms")
	}
	*s = StringLessThan(ss)
	return nil
}
