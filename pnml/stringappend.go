package pnml

import (
	"encoding/xml"
	"errors"
)

type StringAppend struct {
	XMLName xml.Name    `xml:"stringappend"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringAppend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringappend StringAppend
	var ss stringappend
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringAppend: stringappend must have two subterms")
	}
	*s = StringAppend(ss)
	return nil
}
