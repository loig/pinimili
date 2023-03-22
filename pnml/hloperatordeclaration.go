package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type HLOperatorDeclaration struct {
	XMLName              xml.Name                `xml:"namedoperator"`
	ID                   *string                 `xml:"id,attr"`
	Name                 *string                 `xml:"name,attr"`
	VariableDeclarations []HLVariableDeclaration `xml:"parameter>variabledecl"` // optional
	Def                  *HLTermDef              `xml:"def"`
}

func (h *HLOperatorDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type hloperatordeclaration HLOperatorDeclaration
	var hh hloperatordeclaration
	if err := d.DecodeElement(&hh, &start); err != nil {
		return err
	}
	line, col := d.InputPos()
	if hh.ID == nil || *hh.ID == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedoperator without id attribute (or with empty id)")
		return errors.New(msg)
	}
	if hh.Name == nil || *hh.Name == "" {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedoperator without name attribute (or with empty name)")
		return errors.New(msg)
	}
	if hh.Def == nil {
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", namedoperator without def")
		return errors.New(msg)
	}
	*h = HLOperatorDeclaration(hh)
	return nil
}
