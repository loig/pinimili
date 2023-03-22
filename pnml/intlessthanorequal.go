package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type IntLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntLessThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intlessthanorequal IntLessThanOrEqual
	var ii intlessthanorequal
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", leq with ", len(ii.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*i = IntLessThanOrEqual(ii)
	return nil
}
