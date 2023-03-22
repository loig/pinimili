package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", usersort without declaration attribute (or with empty declaration)")
		return errors.New(msg)
	}
	*h = HLUserSort(hh)
	return nil
}
