package pnml

import (
	"encoding/xml"
	"errors"
)

type FIRSort struct {
	XMLName xml.Name `xml:"finiteintrange"`
	Start   *int     `xml:"start,attr"`
	End     *int     `xml:"end,attr"`
}

func (f *FIRSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type firsort FIRSort
	var ff firsort
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	if ff.Start == nil {
		return errors.New("FIRSort: a finiteintrange must have a start attribute")
	}
	if ff.End == nil {
		return errors.New("FIRSort: a finiteintrange must have an end attribute")
	}
	*f = FIRSort(ff)
	return nil
}
