package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type PartitionGreaterThan struct {
	XMLName xml.Name    `xml:"gtp"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (p *PartitionGreaterThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type partitiongreaterthan PartitionGreaterThan
	var pp partitiongreaterthan
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if len(pp.Terms) != 2 {
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", gtp with ", len(pp.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*p = PartitionGreaterThan(pp)
	return nil
}
