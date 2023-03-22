package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if ff.Start == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", finiterange without start attribute")
		return errors.New(msg)
	}
	if ff.End == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", finiterange without end attribute")
		return errors.New(msg)
	}
	*f = FIRSort(ff)
	return nil
}
