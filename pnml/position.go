package pnml

import (
	"encoding/xml"
	"errors"
	"log"
)

type Position struct {
	XMLName xml.Name `xml:"position"`
	X       *float32 `xml:"x,attr"`
	Y       *float32 `xml:"y,attr"`
}

func (p *Position) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type position Position
	var pp position
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.X == nil || pp.Y == nil {
		// the fact that maybe x="" or y="" in the xml file is not checked
		if panicIfNotPnmlCompliant {
			return errors.New("A position must have an x and a y value")
		}
		log.Print("Pinimili:position element without x or y attribute")
	}
	*p = Position(pp)
	return nil
}
