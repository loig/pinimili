package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Pnml struct {
	XMLName xml.Name `xml:"http://www.pnml.org/version-2009/grammar/pnml pnml"`
	Nets    []Net    `xml:"net"` // at least 1
}

func (p *Pnml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type pnml Pnml
	var pp pnml
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if len(pp.Nets) < 1 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", PNML file with ", len(pp.Nets), " net elements (should be at least 1)")
		return errors.New(msg)
	}
	*p = Pnml(pp)
	return nil
}
