package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLSortDeclaration struct {
	XMLName xml.Name `xml:"namedsort"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	Sort    *HLSort  `xml:",any"`
}

func (h *HLSortDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlsortdeclaration HLSortDeclaration
	var hh hlsortdeclaration
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if hh.ID == nil || *hh.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedsort without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if hh.Name == nil || *hh.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedsort without name attribute (or with empty name)")
		return errors.New(msg)
	}
	if hh.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedsort without sort")
		return errors.New(msg)
	}
	*h = HLSortDeclaration(hh)
	return nil
}

func (h HLSortDeclaration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"namedsort"`
		ID      *string  `xml:"id,attr"`
		Name    *string  `xml:"name,attr"`
		Sort    interface{}
	}

	t := tmpH{h.XMLName, h.ID, h.Name, h.Sort.Value}

	return e.Encode(t)
}
