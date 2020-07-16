package pnml

import (
	"encoding/xml"
	"errors"
)

type ListMake struct {
	XMLName xml.Name    `xml:"makelist"`
	Sort    *HLSort     `xml:",any"`
	Terms   []HLSubterm `xml:"subterm"` // 1 or more? what about empty list?
}

func (l *ListMake) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listmake ListMake
	var ll listmake
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Sort == nil {
		return errors.New("ListMake: makelist must have a sort")
	}
	if len(ll.Terms) < 1 {
		return errors.New("ListMake: makelist needs at least one element to build a list")
	}
	*l = ListMake(ll)
	return nil
}

func (l ListMake) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"makelist"`
		Sort    interface{}
		Terms   []HLSubterm `xml:"subterm"`
	}

	t := tmpL{l.XMLName, l.Sort.Value, l.Terms}

	return e.Encode(t)
}
