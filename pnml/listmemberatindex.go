package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if ll.Index == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", memberatindex without index attribute")
		return errors.New(msg)
	}
	if len(ll.Terms) != 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", memberatindex with ", len(ll.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*l = ListMemberAtIndex(ll)
	return nil
}
