package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type FIRLessThanOrEqual struct {
	XMLName xml.Name    `xml:"lessthanorequal"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (f *FIRLessThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firlessthanorequal FIRLessThanOrEqual
	var ff firlessthanorequal
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if len(ff.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", lessthanorequal with ", len(ff.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*f = FIRLessThanOrEqual(ff)
	return nil
}
