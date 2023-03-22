package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringLength struct {
	XMLName xml.Name    `xml:"stringlength"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlength StringLength
	var ss stringlength
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 1 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", stringlength with ", len(ss.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*s = StringLength(ss)
	return nil
}
