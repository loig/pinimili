package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", equality between ", len(bb.Terms), " elements (must be 2)")
		return errors.New(msg)
	}
	*b = BoolEquality(bb)
	return nil
}
