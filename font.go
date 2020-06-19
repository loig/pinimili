package main

import (
	"encoding/xml"
	"errors"
)

type Font struct {
	XMLName    xml.Name `xml:"font"`
	Family     *string  `xml:"family,attr"`     // optional
	Style      *string  `xml:"style,attr"`      // optional
	Weight     *string  `xml:"weight,attr"`     // optional
	Size       *string  `xml:"size,attr"`       // optional
	Decoration *string  `xml:"decoration,attr"` // optional (must be underline, overline, or line-through)
	Align      *string  `xml:"align,attr"`      // optional (must be left, center, or right)
	Rotation   *float32 `xml:"rotation,attr"`   // optional
}

func (f *Font) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type font Font
	var ff font
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if ff.Decoration != nil &&
		*ff.Decoration != "underline" &&
		*ff.Decoration != "overline" &&
		*ff.Decoration != "line-through" {
		return errors.New("A font decoration must be underline, overline, or line-through")
	}
	if ff.Align != nil &&
		*ff.Align != "left" &&
		*ff.Align != "center" &&
		*ff.Align != "right" {
		return errors.New("A font align must be left, center, or right")
	}
	*f = Font(ff)
	return nil
}
