package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if ll.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", makelist without sort")
		return errors.New(msg)
	}
	if len(ll.Terms) < 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", makelist with", len(ll.Terms), "subterm elements (should be at least 1)")
		return errors.New(msg)
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
