package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type StringLessThan struct {
	XMLName xml.Name    `xml:"lts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringlessthan StringLessThan
	var ss stringlessthan
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if len(ss.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", lts with ", len(ss.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*s = StringLessThan(ss)
	return nil
}
