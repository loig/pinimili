package pnml

import (
	"encoding/xml"
	"errors"
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
	if hh.ID == nil || *hh.ID == "" {
		return errors.New("HLOperatorDeclaration: A namedoperator must have a non-empty ID attribute")
	}
	if hh.Name == nil || *hh.Name == "" {
		return errors.New("HLOperatorDeclaration: A namedoperator must have a non-empty name attribute")
	}
	if hh.Def == nil {
		return errors.New("HLOperatorDeclaration: A namedoperator must have a def")
	}
	*h = HLOperatorDeclaration(hh)
	return nil
}
