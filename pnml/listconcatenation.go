package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("ListConcatenation: listconcatenation must have two list elements")
	}
	*l = ListConcatenation(ll)
	return nil
}
