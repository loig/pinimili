package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("FIRLessThan: lessthan must compare two elements")
	}
	*f = FIRLessThan(ff)
	return nil
}
