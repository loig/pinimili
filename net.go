package main

import (
	"encoding/xml"
	"errors"
)

type Net struct {
	XMLName       xml.Name       `xml:"net"`
	ID            *string        `xml:"id,attr"`
	Type          *string        `xml:"type,attr"`
	Name          *Name          `xml:"name"`         // optional
	Pages         []Page         `xml:"page"`         // at least 1
	ToolSpecifics []ToolSpecific `xml:"toolspecific"` // optional
	// net.labels def to be empty
}

func (n *Net) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type net Net
	var nn net
	if err := d.DecodeElement(&nn, &start); err != nil {
		return err
	}
	if nn.ID == nil || *nn.ID == "" {
		return errors.New("A net must have a non-empty id")
	}
	if nn.Type == nil || *nn.Type == "" {
		return errors.New("A net must have a non-empty type")
	}
	if *nn.Type != "http://www.pnml.org/version-2009/grammar/ptnet" {
		return errors.New("Unsupported net type")
	}
	if len(nn.Pages) < 1 {
		return errors.New("A net must have at least one page")
	}
	*n = Net(nn)
	return nil
}
