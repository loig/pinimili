package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type MultisetScalarProduct struct {
	XMLName xml.Name    `xml:"scalarproduct"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (m *MultisetScalarProduct) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetscalarproduct MultisetScalarProduct
	var mm multisetscalarproduct
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	if len(mm.Terms) != 2 { // was 1, do not know why, one should check if it is an issue
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", scalarproduct with ", len(mm.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*m = MultisetScalarProduct(mm)
	return nil
}
