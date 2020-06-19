package main

import (
	"encoding/xml"
	"errors"
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
		return errors.New("An offset must have an x and a y value")
	}
	*o = Offset(oo)
	return nil
}
