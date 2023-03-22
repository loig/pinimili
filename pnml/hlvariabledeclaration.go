package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if hh.ID == nil || *hh.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", variabledecl without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if hh.Name == nil || *hh.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", variabledecl without name attribute (or with empty name)")
		return errors.New(msg)
	}
	if hh.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", variabledecl without sort")
		return errors.New(msg)
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
