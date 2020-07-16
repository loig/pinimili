package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("ListAppend: listappend must have a list element and another element of the sort of the list")
	}
	*l = ListAppend(ll)
	return nil
}
