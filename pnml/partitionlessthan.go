package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", ltp with ", len(pp.Terms), " subterm elements (should be 2)")
		return errors.New(msg)
	}
	*p = PartitionLessThan(pp)
	return nil
}
