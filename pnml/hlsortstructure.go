package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLSortStructure struct {
	XMLName xml.Name `xml:"structure"`
	Sort    *HLSort  `xml:",any"`
}

func (h *HLSortStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlsortstructure HLSortStructure
	var hh hlsortstructure
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.Sort == nil {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", structure without sort")
		return errors.New(msg)
	}
	*h = HLSortStructure(hh)
	return nil
}

func (h HLSortStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"structure"`
		Sort    interface{}
	}

	t := tmpH{h.XMLName, h.Sort.Value}

	return e.Encode(t)
}
