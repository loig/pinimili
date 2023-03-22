package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", successor with ", len(cc.Terms), " elements (must be 1)")
		return errors.New(msg)
	}
	*c = CyclicEnumSuccessor(cc)
	return nil
}
