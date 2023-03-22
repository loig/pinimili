package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type IntGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntGreaterThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intgreaterthanorequal IntGreaterThanOrEqual
	var ii intgreaterthanorequal
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", geq with ", len(ii.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*i = IntGreaterThanOrEqual(ii)
	return nil
}
