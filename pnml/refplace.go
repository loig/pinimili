package pnml

import (
	"encoding/xml"
	"errors"
)

type RefPlace struct {
	XMLName       xml.Name       `xml:"referencePlace"`
	ID            *string        `xml:"id,attr"`
	Reference     *string        `xml:"ref,attr"`
	Name          *Name          `xml:"name"`         // optional
	ToolSpecifics []ToolSpecific `xml:"toolspecific"` // optional
	Graphics      *NodeGraphics  `xml:"graphics"`     // optional
}

func (r *RefPlace) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type refPlace RefPlace
	var rr refPlace
	if err := d.DecodeElement(&rr, &start); err != nil {
		return err
	}
	if rr.ID == nil || *rr.ID == "" {
		return errors.New("A reference place must have a non-empty id")
	}
	if rr.Reference == nil || *rr.Reference == "" {
		return errors.New("A reference place must have a non-empty reference")
	}
	*r = RefPlace(rr)
	return nil
}
