package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringAppend struct {
	XMLName xml.Name    `xml:"stringappend"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringAppend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringappend StringAppend
	var ss stringappend
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", stringappend with ", len(ss.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*s = StringAppend(ss)
	return nil
}
