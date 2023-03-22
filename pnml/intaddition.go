package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type IntAddition struct {
	XMLName xml.Name    `xml:"addition"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntAddition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intaddition IntAddition
	var ii intaddition
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", addition between ", len(ii.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*i = IntAddition(ii)
	return nil
}
