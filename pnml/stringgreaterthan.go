package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringGreaterThan struct {
	XMLName xml.Name    `xml:"gts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringgreaterthan StringGreaterThan
	var ss stringgreaterthan
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", gts with ", len(ss.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*s = StringGreaterThan(ss)
	return nil
}
