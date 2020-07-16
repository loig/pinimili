package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetNumberOf struct {
	XMLName xml.Name    `xml:"numberof"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetNumberOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetnumberof MultisetNumberOf
	var mm multisetnumberof
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		return errors.New("MultisetNumberOf: numberof must have exactly one scalar element and one sort element")
	}
	*m = MultisetNumberOf(mm)
	return nil
}
