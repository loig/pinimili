package pnml

import (
	"encoding/xml"
	"errors"
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
	if pp.ID == nil || *pp.ID == "" {
		return errors.New("PartitionElement: a partitionelement must have a non-empty id attribute")
	}
	if pp.Name == nil || *pp.Name == "" {
		return errors.New("PartitionElement: a partitionelement must have a non-empty name attribute")
	}
	if len(pp.Terms) < 1 {
		return errors.New("PartitionElement: a partitionelement must have at least one term")
	}
	*p = PartitionElement(pp)
	return nil
}
