package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if rr.ID == nil || *rr.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", referencePlace without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if rr.Reference == nil || *rr.Reference == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", referencePlace without ref attribute (or with empty ref)")
		return errors.New(msg)
	}
	*r = RefPlace(rr)
	return nil
}
