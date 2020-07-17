package pnml

import (
	"encoding/xml"
	"errors"
)

type StringConstant struct {
	XMLName xml.Name    `xml:"stringconstant"`
	Value   *string     `xml:"value>text"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringConstant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringconstant StringConstant
	var ss stringconstant
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if ss.Value == nil {
		return errors.New("StringConstant: a stringconstant must have a value attribute")
	}
	if len(ss.Terms) != 0 {
		return errors.New("StringConstant: a stringconstant must have no subterms")
	}
	*s = StringConstant(ss)
	return nil
}
