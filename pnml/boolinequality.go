package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolInequality struct {
	XMLName xml.Name    `xml:"inequality"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolInequality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolinequality BoolInequality
	var bb boolinequality
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		return errors.New("BoolInequality: inequality is always between two elements")
	}
	*b = BoolInequality(bb)
	return nil
}
