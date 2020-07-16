package pnml

import (
	"encoding/xml"
	"errors"
)

type ListMemberAtIndex struct {
	XMLName xml.Name    `xml:"memberatindex"`
	Index   *uint       `xml:"index,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (l *ListMemberAtIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listmemberatindex ListMemberAtIndex
	var ll listmemberatindex
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Index == nil {
		return errors.New("ListMemberAtIndex: memberatindex must have an index attribute")
	}
	if len(ll.Terms) != 1 {
		return errors.New("ListMemberAtIndex: memberatindex must have exactly one list element")
	}
	*l = ListMemberAtIndex(ll)
	return nil
}
