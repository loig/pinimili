package pnml

import (
	"encoding/xml"
	"errors"
)

type MultisetAdd struct {
	XMLName xml.Name    `xml:"add"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetAdd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetadd MultisetAdd
	var mm multisetadd
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) < 2 {
		return errors.New("MultisetAdd: add must have at least two multiset elements")
	}
	*m = MultisetAdd(mm)
	return nil
}
