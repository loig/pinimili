package pnml

import (
	"encoding/xml"
	"errors"
)

type BoolEquality struct {
	XMLName xml.Name    `xml:"equality"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (b *BoolEquality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolequality BoolEquality
	var bb boolequality
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		return errors.New("BoolEquality: equality is always between two elements")
	}
	*b = BoolEquality(bb)
	return nil
}
