package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("MultisetAll: all must not have terms")
	}
	if mm.Sort == nil {
		return errors.New("MultisetAll: all most have a sort")
	}
	*m = MultisetAll(mm)
	return nil
}
