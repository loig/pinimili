package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", greaterthan with ", len(ff.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*f = FIRGreaterThan(ff)
	return nil
}
