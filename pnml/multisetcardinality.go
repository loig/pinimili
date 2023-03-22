package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", cardinality with ", len(mm.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*m = MultisetCardinality(mm)
	return nil
}
