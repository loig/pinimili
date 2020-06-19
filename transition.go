package main

import (
	"encoding/xml"
	"errors"
)

type Transition struct {
	XMLName       xml.Name       `xml:"transition"`
	ID            *string        `xml:"id,attr"`
	Name          *Name          `xml:"name"`         // optional
	ToolSpecifics []ToolSpecific `xml:"toolspecific"` // optional
	Graphics      *NodeGraphics  `xml:"graphics"`     // optional
	// transition.labels def to be empty
}

func (t *Transition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type transition Transition
	var tt transition
	if err := d.DecodeElement(&tt, &start); err != nil {
		return err
	}
	if tt.ID == nil || *tt.ID == "" {
		return errors.New("A transition must have a non-empty id")
	}
	*t = Transition(tt)
	return nil
}
