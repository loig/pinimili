package pnml

import (
	"encoding/xml"
	"errors"
)

type HLUserSort struct {
	XMLName xml.Name `xml:"usersort"`
	ID      *string  `xml:"declaration,attr"` // data type="IDREF"
}

func (h *HLUserSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hlusersort HLUserSort
	var hh hlusersort
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLUserSort: A usersort must have a non-empty declaration attribute")
	}
	*h = HLUserSort(hh)
	return nil
}
