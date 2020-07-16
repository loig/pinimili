package pnml

import (
	"encoding/xml"
	"errors"
)

type FIRConstant struct {
	XMLName xml.Name    `xml:"finiteintrangeconstant"`
	Value   *int        `xml:"value,attr"`
	FIRSort *FIRSort    `xml:"finiteintrange"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (f *FIRConstant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firconstant FIRConstant
	var ff firconstant
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if ff.Value == nil {
		return errors.New("FIRConstant: a finiteintrangeconstant must have a value attribute")
	}
	if ff.FIRSort == nil {
		return errors.New("FIRConstant: a finiteintrangeconstant must have a finiterange element")
	}
	if len(ff.Terms) != 0 {
		return errors.New("FIRConstant: a finiteintrangeconstant must not have subterms")
	}
	*f = FIRConstant(ff)
	return nil
}
