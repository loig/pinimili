package pnml

import (
	"encoding/xml"
	"errors"
)

type StringGreaterThan struct {
	XMLName xml.Name    `xml:"gts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringgreaterthan StringGreaterThan
	var ss stringgreaterthan
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringGreaterThan: gts must have two subterms")
	}
	*s = StringGreaterThan(ss)
	return nil
}
