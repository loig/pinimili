package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", position without x attribute or without y attribute")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*p = Position(pp)
	return nil
}
