package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", inscription without value (or with value 0)")
		return errors.New(msg)
	}
	*p = PTArcAnnotation(pp)
	return nil
}
