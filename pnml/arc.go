package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Arc struct {
	XMLName       xml.Name         `xml:"arc"`
	ID            *string          `xml:"id,attr"`
	Source        *string          `xml:"source,attr"`
	Target        *string          `xml:"target,attr"`
	Name          *Name            `xml:"name"`          // optional
	Weight        *PTArcAnnotation `xml:"inscription"`   // label for pt nets only
	HLInscription *HLAnnotation    `xml:"hlinscription"` // label for hl nets only, optional
	Graphics      *EdgeGraphics    `xml:"graphics"`      // optional
	ToolSpecifics []ToolSpecific   `xml:"toolspecific"`  // optional
	// arc.labels defined to be empty
}

func (a *Arc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type arc Arc
	var aa arc
	if err := d.DecodeElement(&aa, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if aa.ID == nil || *aa.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", an arc must have a non-empty id")
		return errors.New(msg)
	}
	if aa.Source == nil || *aa.Source == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", an arc must have a source")
		return errors.New(msg)
	}
	if aa.Target == nil || *aa.Target == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", an arc must have a target")
		return errors.New(msg)
	}
	*a = Arc(aa)
	return nil
}
