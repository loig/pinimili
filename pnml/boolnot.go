package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", not between ", len(bb.Terms), " elements (must be 1)")
		return errors.New(msg)
	}
	*b = BoolNot(bb)
	return nil
}
