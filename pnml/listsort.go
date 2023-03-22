package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type ListSort struct {
	XMLName xml.Name `xml:"list"`
	Sort    *HLSort  `xml:",any"`
}

func (l *ListSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type listsort ListSort
	var ll listsort
	if err := d.DecodeElement(&ll, &start); err != nil {
		return err
	}
	if ll.Sort == nil {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", list without sort")
		return errors.New(msg)
	}
	*l = ListSort(ll)
	return nil
}

func (l ListSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"list"`
		Sort    interface{}
	}

	t := tmpL{l.XMLName, l.Sort.Value}

	return e.Encode(t)
}
