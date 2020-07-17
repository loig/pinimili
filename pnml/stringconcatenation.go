package pnml

import (
	"encoding/xml"
	"errors"
)

type StringConcatenation struct {
	XMLName xml.Name    `xml:"stringconcatenation"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringConcatenation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringconcatenation StringConcatenation
	var ss stringconcatenation
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		return errors.New("StringConcatenation: stringconcatenation must have two subterms")
	}
	*s = StringConcatenation(ss)
	return nil
}
