package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type Offset struct {
	XMLName xml.Name `xml:"offset"`
	X       *float32 `xml:"x,attr"`
	Y       *float32 `xml:"y,attr"`
}

func (o *Offset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type offset Offset
	var oo offset
	if err := d.DecodeElement(&oo, &start); err != nil {
		return err
	}
	if oo.X == nil || oo.Y == nil {
		// the fact that maybe x="" or y="" in the xml file is not checked
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", offset without x or y attribute")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*o = Offset(oo)
	return nil
}
