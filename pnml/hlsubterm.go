package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("HLSubterm: A subterm must have a term")
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
