package pnml

import (
	"encoding/xml"
	"errors"
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
	if len(mm.Terms) != 0 {
		if panicIfNotPnmlCompliant {
			return errors.New("MultisetEmpty: empty must not have subterms")
		}
		log.Print("Pinimili: empty element with ", len(mm.Terms), " subterm elements (should be 0)")
	}
	if mm.Sort == nil {
		return errors.New("MultisetEmpty: empty must have a sort")
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
