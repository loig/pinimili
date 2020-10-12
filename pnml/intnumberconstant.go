package pnml

import (
	"encoding/xml"
	"errors"
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
	if ii.Value == nil {
		return errors.New("IntNumberConstant: a numberconstant must have a value attribute")
	}
	if len(ii.Terms) != 0 {
		if panicIfNotPnmlCompliant {
			return errors.New("IntNumberConstant: a numberconstant must have no subterms")
		}
		log.Print("Pinimili: numberconstant element with ", len(ii.Terms), " subterm elements (should be 0)")
	}
	numTypes := 0
	if ii.IntNatural != nil {
		numTypes++
		if *ii.Value < 0 {
			return errors.New("IntNumberConstant: a natural numberconstant must be greater or equal than 0")
		}
	}
	if ii.IntPositive != nil {
		numTypes++
		if *ii.Value <= 0 {
			return errors.New("IntNumberConstant: a positive numberconstant must be greater or equal than 0")
		}
	}
	if ii.IntInteger != nil {
		numTypes++
	}
	if numTypes != 1 {
		return errors.New("IntNumberConst: a numberconstant must have exactly one number sort")
	}
	*i = IntNumberConstant(ii)
	return nil
}
