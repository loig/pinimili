package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", fill must have a gradient-rotation attribute which value is vertical, horizontal or diagonal")
		return errors.New(msg)
	}
	*f = Fill(ff)
	return nil
}
