package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if bb.Value == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", a boolean constant must have a value attribute")
		return errors.New(msg)
	}
	if len(bb.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", booleanconstant with ", len(bb.Terms), " subterms (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*b = BoolConstant(bb)
	return nil
}
