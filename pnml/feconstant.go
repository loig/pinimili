package pnml

import (
	"encoding/xml"
	"errors"
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
	if ff.ID == nil || *ff.ID == "" {
		return errors.New("FEConstant: an feconstant must have a non-empty id attribute")
	}
	if ff.Name == nil || *ff.Name == "" {
		return errors.New("FEConstant: an feconstant must have a non-empty name attribute")
	}
	*f = FEConstant(ff)
	return nil
}
