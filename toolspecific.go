package main

import (
	"encoding/xml"
	"errors"
)

type ToolSpecific struct {
	XMLName xml.Name `xml:"toolspecific"`
	Tool    *string  `xml:"tool,attr"`
	Version *string  `xml:"version,attr"`
	// Zero or more anyElement
}

func (t *ToolSpecific) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type toolSpecific ToolSpecific
	var tt toolSpecific
	if err := d.DecodeElement(&tt, &start); err != nil {
		return err
	}
	if tt.Tool == nil || *tt.Tool == "" {
		return errors.New("A tool specific element must have a tool name")
	}
	if tt.Version == nil || *tt.Version == "" {
		return errors.New("A tool specific element must have a version")
	}
	*t = ToolSpecific(tt)
	return nil
}
