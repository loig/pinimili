package main

import (
	"encoding/xml"
	"errors"
)

type PTArcAnnotation struct {
	XMLName       xml.Name            `xml:"inscription"`
	Value         *uint64             `xml:"text"`         // must be integer > 0
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
}

func (p *PTArcAnnotation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type ptArcAnnotation PTArcAnnotation
	var pp ptArcAnnotation
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.Value == nil || *pp.Value == 0 {
		return errors.New("A PT Arc Annotation must have a value (different from 0)")
	}
	*p = PTArcAnnotation(pp)
	return nil
}
