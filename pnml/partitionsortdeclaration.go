package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if pp.ID == nil || *pp.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partition without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if pp.Name == nil || *pp.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partition without name attribute (or with empty name)")
		return errors.New(msg)
	}
	if pp.Sort == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partition without sort")
		return errors.New(msg)
	}
	if len(pp.PartitionElements) < 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partition with ", len(pp.PartitionElements), " partitionelement elements (should be at least 1)")
		return errors.New(msg)
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
