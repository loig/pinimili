package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolAnd struct {
	XMLName xml.Name    `xml:"and"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolAnd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type booland BoolAnd
	var bb booland
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		return errors.New("BoolAnd: and is always between two elements")
	}
	*b = BoolAnd(bb)
	return nil
}
