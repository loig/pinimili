package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type MultisetCardinalityOf struct {
	XMLName xml.Name    `xml:"cardinalityof"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (m *MultisetCardinalityOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetcardinalityof MultisetCardinalityOf
	var mm multisetcardinalityof
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", cardinalityof with ", len(mm.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*m = MultisetCardinalityOf(mm)
	return nil
}
