package pnml

import (
	"encoding/xml"
	"errors"
)

type Pnml struct {
	XMLName xml.Name `xml:"pnml"`
	XMLNs   *string  `xml:"xmlns,attr"` //optional
	Nets    []Net    `xml:"net"`        // at least 1
}

func (p *Pnml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type pnml Pnml
	var pp pnml
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if len(pp.Nets) < 1 {
		return errors.New("A PNML file must contain at least one net")
	}
	*p = Pnml(pp)
	return nil
}
