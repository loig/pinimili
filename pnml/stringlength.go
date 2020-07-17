package pnml

import (
	"encoding/xml"
	"errors"
)

type StringLength struct {
	XMLName xml.Name    `xml:"stringlength"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlength StringLength
	var ss stringlength
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 1 {
		return errors.New("StringLength: stringlength must have exactly one subterm")
	}
	*s = StringLength(ss)
	return nil
}
