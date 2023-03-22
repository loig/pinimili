package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type MultisetNumberOf struct {
	XMLName xml.Name    `xml:"numberof"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetNumberOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetnumberof MultisetNumberOf
	var mm multisetnumberof
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", numberof with ", len(mm.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*m = MultisetNumberOf(mm)
	return nil
}
