package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetSubtract struct {
	XMLName xml.Name    `xml:"subtract"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetSubtract) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetsubtract MultisetSubtract
	var mm multisetsubtract
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		return errors.New("MultisetSubtract: subtract must have exactly two multiset elements")
	}
	*m = MultisetSubtract(mm)
	return nil
}
