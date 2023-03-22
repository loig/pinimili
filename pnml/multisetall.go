package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type MultisetAll struct {
	XMLName xml.Name    `xml:"all"`
	Terms   []HLSubterm `xml:"subterm"`
	Sort    *HLSort     `xml:",any"`
}

func (m *MultisetAll) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type multisetall MultisetAll
	var mm multisetall
	if err := d.DecodeElement(&mm, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if len(mm.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", all with ", len(mm.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
	}
	if mm.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", all without sort")
		return errors.New(msg)
	}
	*m = MultisetAll(mm)
	return nil
}

func (m MultisetAll) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpM struct {
		XMLName xml.Name    `xml:"all"`
		Terms   []HLSubterm `xml:"subterm"`
		Sort    interface{}
	}

	t := tmpM{m.XMLName, m.Terms, m.Sort.Value}

	return e.Encode(t)
}
