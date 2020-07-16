package pnml

import (
	"encoding/xml"
	"errors"
)

type ListEmpty struct {
	XMLName xml.Name    `xml:"emptylist"`
	Sort    *HLSort     `xml:",any"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (l *ListEmpty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listempty ListEmpty
	var ll listempty
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Sort == nil {
		return errors.New("ListEmpty: emptylist must have a sort")
	}
	if len(ll.Terms) != 0 {
		return errors.New("ListEmpty: emptylist must not have subterms")
	}
	*l = ListEmpty(ll)
	return nil
}

func (l ListEmpty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"emptylist"`
		Sort    interface{}
		Terms   []HLSubterm `xml:"subterm"`
	}

	t := tmpL{l.XMLName, l.Sort.Value, l.Terms}

	return e.Encode(t)
}
