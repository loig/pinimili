package pnml

import (
	"encoding/xml"
	"errors"
)

type AnnotationGraphics struct {
	XMLName xml.Name `xml:"graphics"`
	Offset  *Offset  `xml:"offset"`
	Fill    *Fill    `xml:"fill"` // optional
	Line    *Line    `xml:"line"` // optional
	Font    *Font    `xml:"font"` // optional
}

func (a *AnnotationGraphics) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type annotationGraphics AnnotationGraphics
	var aa annotationGraphics
	if err := d.DecodeElement(&aa, &start); err != nil {
		return err
	}
	if aa.Offset == nil {
		return errors.New("An annotation graphics must have an offset")
	}
	*a = AnnotationGraphics(aa)
	return nil
}
