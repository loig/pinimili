package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolOr struct {
	XMLName xml.Name    `xml:"or"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolOr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolor BoolOr
	var bb boolor
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		return errors.New("BoolOr: or is always between two elements")
	}
	*b = BoolOr(bb)
	return nil
}
