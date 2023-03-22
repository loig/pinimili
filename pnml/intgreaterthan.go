package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type IntGreaterThan struct {
	XMLName xml.Name    `xml:"gt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intgreaterthan IntGreaterThan
	var ii intgreaterthan
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", gt with ", len(ii.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*i = IntGreaterThan(ii)
	return nil
}
