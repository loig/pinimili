package pnml

import (
	"encoding/xml"
	"errors"
)

type HLVariableDeclaration struct {
	XMLName xml.Name `xml:"variabledecl"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	Sort    *HLSort  `xml:",any"`
}

func (h *HLVariableDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlvariabledeclaration HLVariableDeclaration
	var hh hlvariabledeclaration
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLVariableDeclaration: A variabledecl must have a non-empty id attribute")
	}
	if hh.Name == nil || *hh.Name == "" {
		return errors.New("HLVariableDeclaration: A variabledecl must have a non-empty name attribute")
	}
	if hh.Sort == nil {
		return errors.New("HLVariableDeclaration: A variabledecl must have a sort")
	}
	*h = HLVariableDeclaration(hh)
	return nil
}

func (h HLVariableDeclaration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"variabledecl"`
		ID      *string  `xml:"id,attr"`
		Name    *string  `xml:"name,attr"`
		Sort    interface{}
	}

	t := tmpH{h.XMLName, h.ID, h.Name, h.Sort.Value}

	return e.Encode(t)
}
