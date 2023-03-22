package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if ll.Start == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", sublist without start attribute")
		return errors.New(msg)
	}
	if ll.Length == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", sublist without length attribute")
		return errors.New(msg)
	}
	if len(ll.Terms) != 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", sublist with ", len(ll.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*l = ListSublist(ll)
	return nil
}
