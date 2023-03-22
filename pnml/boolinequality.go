package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type BoolInequality struct {
	XMLName xml.Name    `xml:"inequality"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (b *BoolInequality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolinequality BoolInequality
	var bb boolinequality
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", inequality between ", len(bb.Terms), " elements (must be 2)")
		return errors.New(msg)
	}
	*b = BoolInequality(bb)
	return nil
}
