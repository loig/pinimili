package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringSubstring struct {
	XMLName xml.Name    `xml:"substring"`
	Start   *uint       `xml:"start,attr"`
	Length  *uint       `xml:"length,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringSubstring) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringsubstring StringSubstring
	var ss stringsubstring
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if ss.Start == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", substring without start attribute")
		return errors.New(msg)
	}
	if ss.Length == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", substring without length attribute")
		return errors.New(msg)
	}
	if len(ss.Terms) != 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", substring with ", len(ss.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*s = StringSubstring(ss)
	return nil
}
