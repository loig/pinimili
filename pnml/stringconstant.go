package pnml

import (
	"encoding/xml"
	"errors"
	"log"
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
		if panicIfNotPnmlCompliant {
			return errors.New("StringConstant: a stringconstant must have no subterms")
		}
		log.Print("Pinimili: stringconstant element with ", len(ss.Terms), " subterm elements (should be 0)")
	}
	*s = StringConstant(ss)
	return nil
}
