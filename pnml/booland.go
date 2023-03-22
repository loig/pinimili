package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type BoolAnd struct {
	XMLName xml.Name    `xml:"and"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (b *BoolAnd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type booland BoolAnd
	var bb booland
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if len(bb.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", and with ", len(bb.Terms), " subterms (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New("msg")
		}
		log.Print("Pinimili: ", msg)
	}
	*b = BoolAnd(bb)
	return nil
}
