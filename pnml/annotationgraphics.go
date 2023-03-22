package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", an annotation graphics must have un offset")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*a = AnnotationGraphics(aa)
	return nil
}
