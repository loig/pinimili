package pnml

import (
	"encoding/xml"
	"errors"
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
	if len(mm.Terms) != 0 {
		if panicIfNotPnmlCompliant {
			return errors.New("MultisetAll: all must not have subterms")
		}
		log.Print("Pinimili: all element with ", len(mm.Terms), " subterm elements (should be 0)")
	}
	if mm.Sort == nil {
		return errors.New("MultisetAll: all must have a sort")
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
