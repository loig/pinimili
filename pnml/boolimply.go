package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolImply struct {
	XMLName xml.Name    `xml:"imply"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolImply) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolimply BoolImply
	var bb boolimply
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		return errors.New("BoolImply: imply is always between two elements")
	}
	*b = BoolImply(bb)
	return nil
}
