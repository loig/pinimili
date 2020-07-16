package pnml

import (
	"encoding/xml"
	"errors"
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
		return errors.New("PartitionGreaterThan: ltp must have two terms")
	}
	*p = PartitionGreaterThan(pp)
	return nil
}
