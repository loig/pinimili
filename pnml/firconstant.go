package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
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
	line, col := d.InputPos()
	if ff.Value == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", finiteintrangeconstant without value attribute")
		return errors.New(msg)
	}
	if ff.FIRSort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", finiteintrangeconstant without finiterange element")
		return errors.New(msg)
	}
	if len(ff.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", finiteintrangeconstant with ", len(ff.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	*f = FIRConstant(ff)
	return nil
}
