package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if pp.Ref == nil || *pp.Ref == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partitionelementof without refpartition attribute (or with empty refpartition)")
		return errors.New(msg)
	}
	if len(pp.Terms) != 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", partitionelementof with ", len(pp.Terms), " subterm elements (should be 1)")
		return errors.New(msg)
	}
	*p = PartitionElementOf(pp)
	return nil
}
