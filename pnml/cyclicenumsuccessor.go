package pnml

import (
	"encoding/xml"
	"errors"
)

type CyclicEnumSuccessor struct {
	XMLName xml.Name    `xml:"successor"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (c *CyclicEnumSuccessor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type cyclicenumsuccessor CyclicEnumSuccessor
	var cc cyclicenumsuccessor
	if err := d.DecodeElement(&cc, &start); err != nil {
		return err
	}
	if len(cc.Terms) != 1 {
		return errors.New("CyclicEnumSuccessor: a successor must have exactly one element")
	}
	*c = CyclicEnumSuccessor(cc)
	return nil
}
