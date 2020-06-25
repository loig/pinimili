package pnml

import (
	"encoding/xml"
	"errors"
)

type Fill struct {
	XMLName          xml.Name `xml:"fill"`
	Color            *string  `xml:"color,attr"`             // optional
	GradientColor    *string  `xml:"gradient-color,attr"`    // optional
	GradientRotation *string  `xml:"gradient-rotation,attr"` // optional (must be vertical, horizontal, or diagonal)
	Image            *string  `xml:"image,attr"`             // optional
}

func (f *Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type fill Fill
	var ff fill
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if ff.GradientRotation != nil &&
		*ff.GradientRotation != "vertical" &&
		*ff.GradientRotation != "horizontal" &&
		*ff.GradientRotation != "diagonal" {
		return errors.New("A fill gradient rotation must be vertical, horizontal, or diagonal")
	}
	*f = Fill(ff)
	return nil
}
