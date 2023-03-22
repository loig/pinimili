package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringGreaterThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringgreaterthanorequal StringGreaterThanOrEqual
	var ss stringgreaterthanorequal
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", geqs with ", len(ss.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*s = StringGreaterThanOrEqual(ss)
	return nil
}
