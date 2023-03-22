package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", graphics without position")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*n = NodeGraphics(nn)
	return nil
}
