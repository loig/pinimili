package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetCardinality struct {
	XMLName xml.Name    `xml:"cardinality"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (m *MultisetCardinality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetcardinality MultisetCardinality
	var mm multisetcardinality
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 1 {
		return errors.New("MultisetCardinality: cardinality must have exactly one multiset element")
	}
	*m = MultisetCardinality(mm)
	return nil
}
