package main

import (
	"encoding/xml"
	"errors"
)

type Line struct {
	XMLName xml.Name `xml:"line"`
	Shape   *string  `xml:"shape,attr"` // optional (must be line or curve)
	Color   *string  `xml:"color,attr"` // optional
	Width   *float32 `xml:"width,attr"` // optional (must be between 0 and 999.9, 4 digits with 1 decimal digit)
	Style   *string  `xml:"style,attr"` // optional (must be solid, dash, or dot)
}

func (l *Line) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type line Line
	var ll line
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Shape != nil &&
		*ll.Shape != "line" &&
		*ll.Shape != "curve" {
		return errors.New("A line shape must be line or curve")
	}
	if ll.Style != nil &&
		*ll.Style != "solid" &&
		*ll.Style != "dash" &&
		*ll.Style != "dot" {
		return errors.New("A line style must be solid, dash, or dot")
	}
	// the value of the width is not tested
	*l = Line(ll)
	return nil
}
