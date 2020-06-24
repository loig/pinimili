package main

import (
	"encoding/xml"
	"errors"
)

type PTMarking struct {
	XMLName       xml.Name            `xml:"initialMarking"`
	Tokens        *uint64             `xml:"text"`         // must be integer >= 0
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
}

func (p *PTMarking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type ptMarking PTMarking
	var pp ptMarking
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.Tokens == nil {
		return errors.New("A PT Marking must have a value")
	}
	*p = PTMarking(pp)
	return nil
}
