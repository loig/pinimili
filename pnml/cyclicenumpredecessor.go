package pnml

import (
	"encoding/xml"
	"errors"
)

type CyclicEnumPredecessor struct {
	XMLName xml.Name    `xml:"predecessor"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (c *CyclicEnumPredecessor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type cyclicenumpredecessor CyclicEnumPredecessor
	var cc cyclicenumpredecessor
	if err := d.DecodeElement(&cc, &start); err != nil {
		return err
	}
	if len(cc.Terms) != 1 {
		return errors.New("CyclicEnumPredecessor: a predecessor must have exactly one element")
	}
	*c = CyclicEnumPredecessor(cc)
	return nil
}
