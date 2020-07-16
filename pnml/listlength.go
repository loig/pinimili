package pnml

import (
	"encoding/xml"
	"errors"
)

type ListLength struct {
	XMLName xml.Name    `xml:"listlength"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (l *ListLength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listlength ListLength
	var ll listlength
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if len(ll.Terms) != 1 {
		return errors.New("ListLength: listlength must have exactly one list element")
	}
	*l = ListLength(ll)
	return nil
}
