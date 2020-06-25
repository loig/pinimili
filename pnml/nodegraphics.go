package pnml

import (
	"encoding/xml"
	"errors"
)

type NodeGraphics struct {
	XMLName   xml.Name   `xml:"graphics"`
	Position  *Position  `xml:"position"`
	Dimension *Dimension `xml:"dimension"` // optional
	Fill      *Fill      `xml:"fill"`      // optional
	Line      *Line      `xml:"line"`      // optional
}

func (n *NodeGraphics) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type nodeGraphics NodeGraphics
	var nn nodeGraphics
	if err := d.DecodeElement(&nn, &start); err != nil {
		return err
	}
	if nn.Position == nil {
		return errors.New("A node (or page) graphics must have a position")
	}
	*n = NodeGraphics(nn)
	return nil
}
