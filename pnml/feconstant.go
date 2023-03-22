package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type FEConstant struct {
	XMLName xml.Name `xml:"feconstant"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
}

func (f *FEConstant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type feconstant FEConstant
	var ff feconstant
	if err := d.DecodeElement(&ff, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if ff.ID == nil || *ff.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", feconstant with empty id attribute")
		return errors.New(msg)
	}
	if ff.Name == nil || *ff.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", feconstant with empty name attribute")
		return errors.New(msg)
	}
	*f = FEConstant(ff)
	return nil
}
