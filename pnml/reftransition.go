package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if rr.ID == nil || *rr.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", referenceTransition without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if rr.Reference == nil || *rr.Reference == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", referenceTransition without ref attribute (or with empty ref)")
		return errors.New(msg)
	}
	*r = RefTransition(rr)
	return nil
}
