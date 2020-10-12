package pnml

import (
	"encoding/xml"
	"errors"
	"log"
)

type BoolConstant struct {
	XMLName xml.Name    `xml:"booleanconstant"`
	Value   *bool       `xml:"value,attr"`
	Terms   []HLSubterm `xml:"subterm"` // should not appear in a boolean constant, but allowed by the XML grammar
}

func (b *BoolConstant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type boolconstant BoolConstant
	var bb boolconstant
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	if bb.Value == nil {
		return errors.New("BoolConstant: a booleanconstant must have a value attribute")
	}
	if len(bb.Terms) != 0 {
		if panicIfNotPnmlCompliant {
			return errors.New("BoolConstant: a booleanconstant must not have subterms")
		}
		log.Print("Pinimili: booleanconstant element with ", len(bb.Terms), " subterm elements (should be 0)")
	}
	*b = BoolConstant(bb)
	return nil
}
