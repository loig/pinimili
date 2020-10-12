package pnml

import (
	"encoding/xml"
	"errors"
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
		if panicIfNotPnmlCompliant {
			return errors.New("BoolOr: or is always between two elements")
		}
		log.Print("Pinimili: or element with ", len(bb.Terms), " subterm elements (should be 2)")
	}
	*b = BoolOr(bb)
	return nil
}
