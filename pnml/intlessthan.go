package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type IntLessThan struct {
	XMLName xml.Name    `xml:"lt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intlessthan IntLessThan
	var ii intlessthan
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", lt with ", len(ii.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*i = IntLessThan(ii)
	return nil
}
