package pnml

import (
	"encoding/xml"
	"errors"
)

type PartitionLessThan struct {
	XMLName xml.Name    `xml:"ltp"`
	Terms   []HLSubterm `xml:"subterm"`
}

func (p *PartitionLessThan) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type partitionlessthan PartitionLessThan
	var pp partitionlessthan
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}
	if len(pp.Terms) != 2 {
		return errors.New("PartitionLessThan: ltp must have two terms")
	}
	*p = PartitionLessThan(pp)
	return nil
}
