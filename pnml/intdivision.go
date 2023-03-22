package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type IntDivision struct {
	XMLName xml.Name    `xml:"div"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntDivision) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intdivision IntDivision
	var ii intdivision
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", div between ", len(ii.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*i = IntDivision(ii)
	return nil
}
