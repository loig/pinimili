package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type IntSubtraction struct {
	XMLName xml.Name    `xml:"subtraction"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntSubtraction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intsubtraction IntSubtraction
	var ii intsubtraction
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", subtraction between ", len(ii.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*i = IntSubtraction(ii)
	return nil
}
