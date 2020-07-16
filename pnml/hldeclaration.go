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

// Lists

type ListEmpty struct {
	XMLName xml.Name    `xml:"emptylist"`
	Sort    *HLSort     `xml:",any"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (l ListEmpty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"emptylist"`
		Sort    interface{}
		Terms   []HLSubterm `xml:"subterm"`
	}

	t := tmpL{l.XMLName, l.Sort.Value, l.Terms}

	return e.Encode(t)
}

// Integers

type IntNatural struct {
	XMLName xml.Name `xml:"natural"`
}

type IntPositive struct {
	XMLName xml.Name `xml:"positive"`
}

type IntInteger struct {
	XMLName xml.Name `xml:"integer"`
}

type IntAddition struct {
	XMLName xml.Name    `xml:"addition"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntSubtraction struct {
	XMLName xml.Name    `xml:"subtraction"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntMultiplication struct {
	XMLName xml.Name    `xml:"mult"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntDivision struct {
	XMLName xml.Name    `xml:"div"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntModulo struct {
	XMLName xml.Name    `xml:"mod"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntGreaterThan struct {
	XMLName xml.Name    `xml:"gt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntLessThan struct {
	XMLName xml.Name    `xml:"lt"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leq"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type IntNumberConstant struct {
	XMLName xml.Name    `xml:"numberconstant"`
	Value   *int        `xml:"value,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
	// choice
	IntNatural  *IntNatural  `xml:"natural"`
	IntPositive *IntPositive `xml:"positive"`
	IntInteger  *IntInteger  `xml:"integer"`
	// end choice
}

// Strings

type StringSort struct {
	XMLName xml.Name `xml:"string"`
}

type StringLessThan struct {
	XMLName xml.Name    `xml:"lts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringLessThanOrEqual struct {
	XMLName xml.Name    `xml:"leqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringGreaterThan struct {
	XMLName xml.Name    `xml:"gts"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"geqs"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringConcatenation struct {
	XMLName xml.Name    `xml:"stringconcatenation"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type StringAppend struct {
	XMLName xml.Name    `xml:"stringappend"`
	Terms   []HLSubterm `xml:"subterm"` // optional
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
