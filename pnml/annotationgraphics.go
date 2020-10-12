package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("An annotation graphics must have an offset")
		}
		log.Print("Pinimili: graphics element with no offset element")
	}
	*a = AnnotationGraphics(aa)
	return nil
}
