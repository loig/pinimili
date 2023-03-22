package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLessThanOrEqual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlessthanorequal StringLessThanOrEqual
	var ss stringlessthanorequal
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", leqs with ", len(ss.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*s = StringLessThanOrEqual(ss)
	return nil
}
