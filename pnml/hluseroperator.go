package pnml

import (
	"encoding/xml"
	"errors"
)

type HLUserOperator struct { // recursion forbidden
	XMLName xml.Name    `xml:"useroperator"`
	ID      *string     `xml:"declaration,attr"` // data type="IDREF"
	Terms   []HLSubterm `xml:"subterm"`          // optional
}

func (h *HLUserOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hluseroperator HLUserOperator
	var hh hluseroperator
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLUserOperator: A useroperator must have a non-empty declaration attribute")
	}
	*h = HLUserOperator(hh)
	return nil
}
