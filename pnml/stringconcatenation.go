package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("StringConcatenation: stringconcatenation must have two subterms")
		}
		log.Print("Pinimili: stringconcatenation element with ", len(ss.Terms), " subterm elements (should be 2)")
	}
	*s = StringConcatenation(ss)
	return nil
}
