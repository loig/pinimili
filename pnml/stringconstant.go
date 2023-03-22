package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if ss.Value == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", stringconstant without value attribute")
		return errors.New(msg)
	}
	if len(ss.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", stringconstant with ", len(ss.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*s = StringConstant(ss)
	return nil
}
