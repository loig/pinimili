package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type IntModulo struct {
	XMLName xml.Name    `xml:"mod"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntModulo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intmodulo IntModulo
	var ii intmodulo
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", mod with ", len(ii.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*i = IntModulo(ii)
	return nil
}
