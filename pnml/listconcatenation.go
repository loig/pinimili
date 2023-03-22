package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type ListConcatenation struct {
	XMLName xml.Name    `xml:"listconcatenation"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (l *ListConcatenation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listconcatenation ListConcatenation
	var ll listconcatenation
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if len(ll.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", listconcatenation with ", len(ll.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*l = ListConcatenation(ll)
	return nil
}
