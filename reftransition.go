package main

import (
	"encoding/xml"
	"errors"
)

type RefTransition struct {
	XMLName       xml.Name       `xml:"referenceTransition"`
	ID            *string        `xml:"id,attr"`
	Reference     *string        `xml:"ref,attr"`
	Name          *Name          `xml:"name"`         // optional
	ToolSpecifics []ToolSpecific `xml:"toolspecific"` // optional
	Graphics      *NodeGraphics  `xml:"graphics"`     // optional
}

func (r *RefTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type refTransition RefTransition
	var rr refTransition
	if err := d.DecodeElement(&rr, &start); err != nil {
		return err
	}
	if rr.ID == nil || *rr.ID == "" {
		return errors.New("A reference transition must have a non-empty id")
	}
	if rr.Reference == nil || *rr.Reference == "" {
		return errors.New("A reference transition must have a non-empty reference")
	}
	*r = RefTransition(rr)
	return nil
}
