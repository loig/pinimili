package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Transition struct {
	XMLName       xml.Name       `xml:"transition"`
	ID            *string        `xml:"id,attr"`
	Name          *Name          `xml:"name"`         // optional
	ToolSpecifics []ToolSpecific `xml:"toolspecific"` // optional
	Condition     *HLCondition   `xml:"condition"`    // label for hl nets only, optional
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", transition element without id attribute (or with empty id)")
		return errors.New(msg)
	}
	*t = Transition(tt)
	return nil
}
