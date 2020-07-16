package pnml

import (
	"encoding/xml"
	"errors"
)

type PartitionSortDeclaration struct {
	XMLName           xml.Name           `xml:"partition"`
	ID                *string            `xml:"id,attr"`
	Name              *string            `xml:"name,attr"`
	Sort              *HLSort            `xml:",any"`
	PartitionElements []PartitionElement `xml:"partitionelement"` // one or more
}

func (p *PartitionSortDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type partitionsortdeclaration PartitionSortDeclaration
	var pp partitionsortdeclaration
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.ID == nil || *pp.ID == "" {
		return errors.New("PartitionSortDeclaration: a partition must have a non-empty id attribute")
	}
	if pp.Name == nil || *pp.Name == "" {
		return errors.New("PartitionSortDeclaration: a partition must have a non-empty name attribute")
	}
	if pp.Sort == nil {
		return errors.New("PartitionSortDeclaration: a partition must have a sort")
	}
	if len(pp.PartitionElements) < 1 {
		return errors.New("PartitionSortDeclaration: a partition must have at least one partitionelement")
	}
	*p = PartitionSortDeclaration(pp)
	return nil
}

func (p PartitionSortDeclaration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpP struct {
		XMLName           xml.Name `xml:"partition"`
		ID                *string  `xml:"id,attr"`
		Name              *string  `xml:"name,attr"`
		Sort              interface{}
		PartitionElements []PartitionElement `xml:"partitionelement"`
	}

	t := tmpP{p.XMLName, p.ID, p.Name, p.Sort.Value, p.PartitionElements}

	return e.Encode(t)
}
