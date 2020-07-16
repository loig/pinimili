package pnml

import (
	"encoding/xml"
	"errors"
)

type FIRGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"greaterthanorequal"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (f *FIRGreaterThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firgreaterthanorequal FIRGreaterThanOrEqual
	var ff firgreaterthanorequal
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if len(ff.Terms) != 2 {
		return errors.New("FIRGreaterThanOrEqual: greaterthanorequal must compare two elements")
	}
	*f = FIRGreaterThanOrEqual(ff)
	return nil
}
