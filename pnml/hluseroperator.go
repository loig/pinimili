package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", useroperator without declaration attribute (or with empty declaration)")
		return errors.New(msg)
	}
	*h = HLUserOperator(hh)
	return nil
}
