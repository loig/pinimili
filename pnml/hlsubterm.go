package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLSubterm struct {
	XMLName xml.Name `xml:"subterm"`
	Term    *HLTerm  `xml:",any"`
}

func (h *HLSubterm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlsubterm HLSubterm
	var hh hlsubterm
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.Term == nil {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", subterm without term")
		return errors.New(msg)
	}
	*h = HLSubterm(hh)
	return nil
}

func (h HLSubterm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"subterm"`
		Term    interface{}
	}

	t := tmpH{h.XMLName, h.Term.Value}

	return e.Encode(t)
}
