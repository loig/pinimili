package pnml

import (
	"encoding/xml"
	"errors"
)

type FIRGreaterThan struct {
	XMLName xml.Name    `xml:"greaterthan"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (f *FIRGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firgreaterthan FIRGreaterThan
	var ff firgreaterthan
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if len(ff.Terms) != 2 {
		return errors.New("FIRGreaterThan: greaterthan must compare two elements")
	}
	*f = FIRGreaterThan(ff)
	return nil
}
