package pnml

import (
	"encoding/xml"
	"errors"
)

type PartitionElementOf struct {
	XMLName xml.Name    `xml:"partitionelementof"`
	Ref     *string     `xml:"refpartition,attr"` // data of type IDREF
	Terms   []HLSubterm `xml:"subterm"`           // optional
}

func (p *PartitionElementOf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type partitionelementof PartitionElementOf
	var pp partitionelementof
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if pp.Ref == nil || *pp.Ref == "" {
		return errors.New("PartitionElementOf: partitionelementof must have a non-empty refpartition attribute")
	}
	if len(pp.Terms) != 1 {
		return errors.New("PartitionElementOf: partitionelementof must have one term")
	}
	*p = PartitionElementOf(pp)
	return nil
}
