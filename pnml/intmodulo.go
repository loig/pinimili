package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("IntModulo: modulo must have two subterms")
	}
	*i = IntModulo(ii)
	return nil
}
