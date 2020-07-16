package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("FIRLessThanOrEqual: lessthanorequal must compare two elements")
	}
	*f = FIRLessThanOrEqual(ff)
	return nil
}
