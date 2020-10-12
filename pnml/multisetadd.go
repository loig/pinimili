package pnml

import (
	"encoding/xml"
	"errors"
	"log"
)

type MultisetAdd struct {
	XMLName xml.Name    `xml:"add"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (m *MultisetAdd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetadd MultisetAdd
	var mm multisetadd
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) < 2 {
		if panicIfNotPnmlCompliant {
			return errors.New("MultisetAdd: add must have at least two multiset elements")
		}
		log.Print("Pinimili: add element with ", len(mm.Terms), " subterm elements (should be at least 2)")
	}
	*m = MultisetAdd(mm)
	return nil
}
