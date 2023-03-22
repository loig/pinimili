package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type PartitionElement struct {
	XMLName xml.Name `xml:"partitionelement"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	Terms   []HLTerm `xml:",any"`
}

func (p *PartitionElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type partitionelement PartitionElement
	var pp partitionelement
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if pp.ID == nil || *pp.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partitionelement without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if pp.Name == nil || *pp.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partitionelement without name attribute (or with empty name)")
		return errors.New(msg)
	}
	if len(pp.Terms) < 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partitionelement with ", len(pp.Terms), " terms (should be at least 1)")
		return errors.New(msg)
	}
	*p = PartitionElement(pp)
	return nil
}
