package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
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
	line, col := d.InputPos()
	if ll.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", emptylist without sort")
		return errors.New(msg)
	}
	if len(ll.Terms) != 0 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", emptylist with ", len(ll.Terms), " subterm elements (should be 0)")
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		log.Print("Pinimili: ", msg)
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
