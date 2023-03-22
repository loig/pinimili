package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Net struct {
	XMLName        xml.Name        `xml:"net"`
	ID             *string         `xml:"id,attr"`
	Type           *string         `xml:"type,attr"`
	Name           *Name           `xml:"name"`         // optional
	HLDeclarations []HLDeclaration `xml:"declaration"`  // optional
	Pages          []Page          `xml:"page"`         // at least 1
	ToolSpecifics  []ToolSpecific  `xml:"toolspecific"` // optional
	// net.labels def to be empty
}

func (n *Net) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type net Net
	var nn net
	if err := d.DecodeElement(&nn, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if nn.ID == nil || *nn.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", net without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if nn.Type == nil || *nn.Type == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", net without type attribute (or with empty type)")
		return errors.New(msg)
	}
	if *nn.Type != "http://www.pnml.org/version-2009/grammar/ptnet" &&
		*nn.Type != "http://www.pnml.org/version-2009/grammar/symmetricnet" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", unsupported net type")
		return errors.New(msg)
	}
	if len(nn.Pages) < 1 {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", net with ", len(nn.Pages), " page elements (should be at least 1)")
		return errors.New(msg)
	}
	*n = Net(nn)
	return nil
}
