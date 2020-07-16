package pnml

import (
	"encoding/xml"
	"errors"
)

type ListSublist struct {
	XMLName xml.Name    `xml:"sublist"`
	Start   *uint       `xml:"start,attr"`
	Length  *uint       `xml:"length,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (l *ListSublist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listsublist ListSublist
	var ll listsublist
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Start == nil {
		return errors.New("ListSublist: sublist must have a start attribute")
	}
	if ll.Length == nil {
		return errors.New("ListSublist: sublist must have a length attribute")
	}
	if len(ll.Terms) != 1 {
		return errors.New("ListSublist: sublist must have exactly one list element")
	}
	*l = ListSublist(ll)
	return nil
}
