package main

import (
	"encoding/xml"
	"errors"
)

type PageGraphics = NodeGraphics

type Page struct {
	XMLName        xml.Name        `xml:"page"`
	ID             *string         `xml:"id,attr"`
	Name           *Name           `xml:"name"`                // optional
	ToolSpecifics  []ToolSpecific  `xml:"toolspecific"`        // optional
	Pages          []Page          `xml:"page"`                // optional
	Places         []Place         `xml:"place"`               // optional
	Transitions    []Transition    `xml:"transition"`          // optional
	RefPlaces      []RefPlace      `xml:"referencePlace"`      // optional
	RefTransitions []RefTransition `xml:"referenceTransition"` // optional
	Arcs           []Arc           `xml:"arc"`                 // optional
	Graphics       *PageGraphics   `xml:"graphics"`            // optional
	// page.labels def to be empty
}

func (p *Page) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type page Page
	var pp page
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.ID == nil || *pp.ID == "" {
		return errors.New("A page must have a non-empty id")
	}
	*p = Page(pp)
	return nil
}
