package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type MultisetAdd struct {
	XMLName xml.Name    `xml:"add"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (m *MultisetAdd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetadd MultisetAdd
	var mm multisetadd
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) < 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", add with ", len(mm.Terms), " subterm elements (should be at least 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*m = MultisetAdd(mm)
	return nil
}
