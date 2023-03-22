package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", predecessor with ", len(cc.Terms), " elements (must be 1)")
		return errors.New(msg)
	}
	*c = CyclicEnumPredecessor(cc)
	return nil
}
