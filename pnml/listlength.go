package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", listlength with ", len(ll.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*l = ListLength(ll)
	return nil
}
