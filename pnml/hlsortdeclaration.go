package pnml

import (
	"encoding/xml"
	"errors"
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
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLSortDeclaration: A namedsort must have a non-empty id attribute")
	}
	if hh.Name == nil || *hh.Name == "" {
		return errors.New("HLSortDeclaration: A namedsort must have a non-empty name attribute")
	}
	if hh.Sort == nil {
		return errors.New("HLSortDeclaration: A namedsort must have a sort")
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
