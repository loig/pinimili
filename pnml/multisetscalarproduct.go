package pnml

import (
	"encoding/xml"
	"errors"
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
	if len(mm.Terms) != 1 { // should be 2 ?
		return errors.New("MultisetScalarProduct: scalarproduct must have exactly one scalar element and one multiset element")
	}
	*m = MultisetScalarProduct(mm)
	return nil
}
