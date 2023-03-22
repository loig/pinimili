package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLTermStructure struct {
	XMLName xml.Name `xml:"structure"`
	Term    *HLTerm  `xml:",any"`
}

func (h *HLTermStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hltermstructure HLTermStructure
	var hh hltermstructure
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.Term == nil {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", structure without term")
		return errors.New(msg)
	}
	*h = HLTermStructure(hh)
	return nil
}

func (h HLTermStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"structure"`
		Term    interface{}
	}

	t := tmpH{h.XMLName, h.Term.Value}

	return e.Encode(t)
}
