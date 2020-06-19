package main

import (
	"encoding/xml"
	"errors"
)

type Place struct {
	XMLName        xml.Name       `xml:"place"`
	ID             *string        `xml:"id,attr"`
	Name           *Name          `xml:"name"`           // optional
	ToolSpecifics  []ToolSpecific `xml:"toolspecific"`   // optional
	InitialMarking *PTMarking     `xml:"initialMarking"` // label for pt nets only
	Graphics       *NodeGraphics  `xml:"graphics"`       // optional
	// place.labels defined to be empty
}

func (p *Place) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type place Place
	var pp place
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.ID == nil || *pp.ID == "" {
		return errors.New("A place must have a non-empty id")
	}
	*p = Place(pp)
	return nil
}
