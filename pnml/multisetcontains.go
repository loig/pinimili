package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	if len(mm.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", contains with ", len(mm.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*m = MultisetContains(mm)
	return nil
}
