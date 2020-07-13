package pnml

import (
	"encoding/xml"
	"errors"
)

type HLMultisetSort struct {
	XMLName xml.Name `xml:"multisetsort"`
	Sort    *HLSort  `xml:",any"`
}

func (h *HLMultisetSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlmultisetsort HLMultisetSort
	var hh hlmultisetsort
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.Sort == nil {
		return errors.New("HLMultisetSort: A multisetsort must have a sort")
	}
	*h = HLMultisetSort(hh)
	return nil
}

func (h HLMultisetSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"multisetsort"`
		Sort    interface{}
	}

	t := tmpH{h.XMLName, h.Sort.Value}

	return e.Encode(t)
}
