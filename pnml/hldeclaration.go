package pnml

import (
	"encoding/xml"
)

// TODO: propagate types to check them at operation level

type HLDeclaration struct {
	XMLName                        xml.Name                   `xml:"declaration"`
	Info                           *string                    `xml:"text"`                                    // optional
	SortDeclarations               []HLSortDeclaration        `xml:"structure>declarations>namedsort"`        // optional
	PartitionSortDeclarations      []PartitionSortDeclaration `xml:"structure>declarations>partition"`        // optional
	VariableDeclarations           []HLVariableDeclaration    `xml:"structure>declarations>variabledecl"`     // optional
	OperatorDeclarations           []HLOperatorDeclaration    `xml:"structure>declarations>namedoperator"`    // optional
	PartitionOperatorsDeclarations []PartitionElement         `xml:"structure>declarations>partitionelement"` // optional
	FEConstantDeclarations         []FEConstant               `xml:"structure>declarations>feconstant"`       // optional
}

type StringLength struct {
	XMLName xml.Name    `xml:"stringlength"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringSubstring struct {
	XMLName xml.Name    `xml:"substring"`
	Start   *uint       `xml:"start,attr"`
	Length  *uint       `xml:"length,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringConstant struct {
	XMLName xml.Name    `xml:"stringconstant"`
	Value   *string     `xml:"value>text"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}
