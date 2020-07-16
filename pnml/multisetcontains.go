package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetContains struct {
	XMLName xml.Name    `xml:"contains"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (m *MultisetContains) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetcontains MultisetContains
	var mm multisetcontains
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 1 {
		return errors.New("MultisetContains: contains must have exactly two multiset elements")
	}
	*m = MultisetContains(mm)
	return nil
}
