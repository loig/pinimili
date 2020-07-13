package pnml

import (
	"encoding/xml"
	"errors"
)

type HLVariable struct {
	XMLName xml.Name `xml:"variable"`
	ID      *string  `xml:"refvariable,attr"` // data type="IDREF"
}

func (h *HLVariable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlvariable HLVariable
	var hh hlvariable
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLVariable: A variable must have a non-empty refvariable attribute")
	}
	*h = HLVariable(hh)
	return nil
}
