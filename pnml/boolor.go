package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", or between ", len(bb.Terms), " elements (must be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*b = BoolOr(bb)
	return nil
}
