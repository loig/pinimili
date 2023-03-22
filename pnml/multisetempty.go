package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type MultisetEmpty struct {
	XMLName xml.Name    `xml:"empty"`
	Terms   []HLSubterm `xml:"subterm"` // optional
	Sort    *HLSort     `xml:",any"`
}

func (m *MultisetEmpty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetempty MultisetEmpty
	var mm multisetempty
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if len(mm.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", empty with ", len(mm.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	if mm.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", empty without sort")
		return errors.New(msg)
	}
	*m = MultisetEmpty(mm)
	return nil
}

func (m MultisetEmpty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpM struct {
		XMLName xml.Name    `xml:"empty"`
		Terms   []HLSubterm `xml:"subterm"`
		Sort    interface{}
	}

	t := tmpM{m.XMLName, m.Terms, m.Sort.Value}

	return e.Encode(t)
}
