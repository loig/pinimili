package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolNot struct {
	XMLName xml.Name    `xml:"not"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolNot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolnot BoolNot
	var bb boolnot
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 1 {
		return errors.New("BoolNot: not is always for one element")
	}
	*b = BoolNot(bb)
	return nil
}
