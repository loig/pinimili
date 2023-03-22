package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type IntNumberConstant struct {
	XMLName xml.Name    `xml:"numberconstant"`
	Value   *int        `xml:"value,attr"`
	Terms   []HLSubterm `xml:"subterm"` // empty
	// choice
	IntNatural  *IntNatural  `xml:"natural"`
	IntPositive *IntPositive `xml:"positive"`
	IntInteger  *IntInteger  `xml:"integer"`
	// end choice
}

func (i *IntNumberConstant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intnumberconstant IntNumberConstant
	var ii intnumberconstant
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if ii.Value == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", numberconstant without value attribute")
		return errors.New(msg)
	}
	if len(ii.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", numberconstant with ", len(ii.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	numTypes := 0
	if ii.IntNatural != nil {
		numTypes++
		if *ii.Value < 0 {
			msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", natural numberconstant must lower than 0")
			return errors.New(msg)
		}
	}
	if ii.IntPositive != nil {
		numTypes++
		if *ii.Value <= 0 {
			msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", positive numberconstant lower or equal than 0")
			return errors.New(msg)
		}
	}
	if ii.IntInteger != nil {
		numTypes++
	}
	if numTypes != 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", numberconstant with ", numTypes, "number sorts (should be 1)")
		return errors.New(msg)
	}
	*i = IntNumberConstant(ii)
	return nil
}
