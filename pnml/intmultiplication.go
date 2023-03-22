package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type IntMultiplication struct {
	XMLName xml.Name    `xml:"mult"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntMultiplication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intmultiplication IntMultiplication
	var ii intmultiplication
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", mult between ", len(ii.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*i = IntMultiplication(ii)
	return nil
}
