package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLTermDef struct {
	XMLName xml.Name `xml:"def"`
	Term    *HLTerm  `xml:",any"`
}

func (h *HLTermDef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hltermdef HLTermDef
	var hh hltermdef
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.Term == nil {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", def without term")
		return errors.New(msg)
	}
	*h = HLTermDef(hh)
	return nil
}

func (h HLTermDef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"def"`
		Term    interface{}
	}

	t := tmpH{h.XMLName, h.Term.Value}

	return e.Encode(t)
}
