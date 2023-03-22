package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Dimension struct {
	XMLName xml.Name `xml:"dimension"`
	X       *float32 `xml:"x,attr"` // between 0 and 999.9, 4 digits with 1 decimal digit
	Y       *float32 `xml:"y,attr"` // between 0 and 999.9, 4 digits with 1 decimal digit
}

func (di *Dimension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type dimension Dimension
	var dd dimension
	if err := d.DecodeElement(&dd, &start); err != nil {
		return err
	}
	if dd.X == nil || dd.Y == nil {
		// the fact that maybe x="" or y="" in the xml file is not checked
		// the constraint on the number of digits is not tested
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", dimension with missing x or missing y")
		return errors.New(msg)
	}
	*di = Dimension(dd)
	return nil
}
