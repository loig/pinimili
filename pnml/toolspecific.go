package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	line, col := d.InputPos()
	if tt.Tool == nil || *tt.Tool == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", toolspecific element without tool attribute (or with empty tool)")
		return errors.New(msg)
	}
	if tt.Version == nil || *tt.Version == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", toolspecific element without version attribute (or with empty version)")
		return errors.New(msg)
	}
	*t = ToolSpecific(tt)
	return nil
}
