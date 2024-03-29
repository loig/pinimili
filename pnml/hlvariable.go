package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", variable without refvariable attribute (or with empty refvariable)")
		return errors.New(msg)
	}
	*h = HLVariable(hh)
	return nil
}
