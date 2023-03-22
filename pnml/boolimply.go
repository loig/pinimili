package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", imply between ", len(bb.Terms), " elements (must be 2)")
		return errors.New(msg)
	}
	*b = BoolImply(bb)
	return nil
}
