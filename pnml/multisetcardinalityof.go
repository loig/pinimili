package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetCardinalityOf struct {
	XMLName xml.Name    `xml:"cardinalityof"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetCardinalityOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetcardinalityof MultisetCardinalityOf
	var mm multisetcardinalityof
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		return errors.New("MultisetCardinalityOf: cardinalityof must have exactly two elements (a multiset and a sort)")
	}
	*m = MultisetCardinalityOf(mm)
	return nil
}
