package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", greaterthanorequal with ", len(ff.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*f = FIRGreaterThanOrEqual(ff)
	return nil
}
