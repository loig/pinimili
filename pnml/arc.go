package pnml

import (
	"encoding/xml"
	"errors"
)

type Arc struct {
	XMLName       xml.Name         `xml:"arc"`
	ID            *string          `xml:"id,attr"`
	Source        *string          `xml:"source,attr"`
	Target        *string          `xml:"target,attr"`
	Name          *Name            `xml:"name"`         // optional
	Weight        *PTArcAnnotation `xml:"inscription"`  // label for pt nets only
	Graphics      *EdgeGraphics    `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific   `xml:"toolspecific"` // optional
	// arc.labels defined to be empty
}

func (a *Arc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type arc Arc
	var aa arc
	if err := d.DecodeElement(&aa, &start); err != nil {
		return err
	}
	if aa.ID == nil || *aa.ID == "" {
		return errors.New("An arc must have a non-empty id")
	}
	if aa.Source == nil || *aa.Source == "" {
		return errors.New("An arc must have a source")
	}
	if aa.Target == nil || *aa.Target == "" {
		return errors.New("An arc must have a target")
	}
	*a = Arc(aa)
	return nil
}
