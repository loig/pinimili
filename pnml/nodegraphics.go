package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("A node (or page) graphics must have a position")
		}
		log.Print("Pinimili: graphics element with no position element")
	}
	*n = NodeGraphics(nn)
	return nil
}
