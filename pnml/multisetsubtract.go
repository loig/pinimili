package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type MultisetSubtract struct {
	XMLName xml.Name    `xml:"subtract"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetSubtract) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetsubtract MultisetSubtract
	var mm multisetsubtract
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", subtract with ", len(mm.Terms), " subterm elements (should be 2)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*m = MultisetSubtract(mm)
	return nil
}
