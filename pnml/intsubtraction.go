package pnml

import (
	"encoding/xml"
	"errors"
	"log"
)

type IntSubtraction struct {
	XMLName xml.Name    `xml:"subtraction"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (i *IntSubtraction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type intsubtraction IntSubtraction
	var ii intsubtraction
	if err := d.DecodeElement(&ii, &start); err != nil {
		return err
	}
	if len(ii.Terms) != 2 {
		if panicIfNotPnmlCompliant {
			return errors.New("IntSubtraction: subtraction must have two subterms")
		}
		log.Print("Pinimili: subtraction element with ", len(ii.Terms), " subterm elements (should be 2)")
	}
	*i = IntSubtraction(ii)
	return nil
}
