package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type ListAppend struct {
	XMLName xml.Name    `xml:"listappend"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (l *ListAppend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listappend ListAppend
	var ll listappend
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if len(ll.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", listappend with ", len(ll.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*l = ListAppend(ll)
	return nil
}
