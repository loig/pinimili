package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type FIRLessThan struct {
	XMLName xml.Name    `xml:"lessthan"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (f *FIRLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firlessthan FIRLessThan
	var ff firlessthan
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if len(ff.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", lessthan with ", len(ff.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*f = FIRLessThan(ff)
	return nil
}
