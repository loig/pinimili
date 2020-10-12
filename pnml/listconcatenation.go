package pnml

import (
	"encoding/xml"
	"errors"
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
		if panicIfNotPnmlCompliant {
			return errors.New("ListConcatenation: listconcatenation must have two list elements")
		}
		log.Print("Pinimili: listconcatenation element with ", len(ll.Terms), " subterm elements (should be 2)")
	}
	*l = ListConcatenation(ll)
	return nil
}
