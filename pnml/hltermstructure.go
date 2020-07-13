package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("HLTermStructure: A structure must have a term")
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
