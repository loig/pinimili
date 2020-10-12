package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("MultisetSubtract: subtract must have exactly two multiset elements")
		}
		log.Print("Pinimili: subtract element with ", len(mm.Terms), " subterm elements (should be 2)")
	}
	*m = MultisetSubtract(mm)
	return nil
}
