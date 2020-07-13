package pnml

import (
	"encoding/xml"
	"log"
)

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

type HLTermStructure struct {
	XMLName xml.Name `xml:"structure"`
	Term    *HLTerm  `xml:",any"`
}

func (h HLTermStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"structure"`
		Term    interface{}
	}

	t := tmpH{h.XMLName, h.Term.Value}

	return e.Encode(t)
}

type HLMarking struct {
	XMLName       xml.Name            `xml:"hlinitialMarking"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}

type HLCondition struct {
	XMLName       xml.Name            `xml:"condition"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}

type HLAnnotation struct {
	XMLName       xml.Name            `xml:"hlinscription"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}

type HLTerm struct {
	Type  string `xml:"-"`
	Value interface{}
}

func (h *HLTerm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	h.Type = start.Name.Local
	switch h.Type {
	// built in operators
	case "variable":
		var term HLVariable
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "equality":
		var term BoolEquality
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "inequality":
		var term BoolInequality
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "and":
		var term BoolAnd
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "or":
		var term BoolOr
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "imply":
		var term BoolImply
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "not":
		var term BoolNot
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "predecessor":
		var term CyclicEnumPredecessor
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "successor":
		var term CyclicEnumSuccessor
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "lessthan":
		var term FIRLessThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "lessthanorequal":
		var term FIRLessThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "greaterthan":
		var term FIRGreaterThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "greaterthanorequal":
		var term FIRGreaterThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "cardinality":
		var term MultisetCardinality
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "cardinalityof":
		var term MultisetCardinalityOf
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "contains":
		var term MultisetContains
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "ltp":
		var term PartitionLessThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "gtp":
		var term PartitionGreaterThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "partitionelementof":
		var term PartitionElementOf
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "listappend":
		var term ListAppend
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "listconcatenation":
		var term ListConcatenation
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "makelist":
		var term ListMake
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "listlength":
		var term ListLength
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "memberatindex":
		var term ListMemberAtIndex
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "sublist":
		var term ListSublist
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "addition":
		var term IntAddition
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "subtraction":
		var term IntSubtraction
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "mult":
		var term IntMultiplication
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "div":
		var term IntDivision
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "mod":
		var term IntModulo
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "gt":
		var term IntGreaterThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "geq":
		var term IntGreaterThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "lt":
		var term IntLessThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "leq":
		var term IntLessThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "lts":
		var term StringLessThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "leqs":
		var term StringLessThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "gts":
		var term StringGreaterThan
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "geqs":
		var term StringGreaterThanOrEqual
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "stringconcatenation":
		var term StringConcatenation
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "stringappend":
		var term StringAppend
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "stringlength":
		var term StringLength
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "substring":
		var term StringSubstring
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
		// built in constants
	case "booleanconstant":
		var term BoolConstant
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "finiteintrangeconstant":
		var term FIRConstant
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "dotconstant":
		var term DotConstant
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "emptylist":
		var term ListEmpty
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "numberconstant":
		var term IntNumberConstant
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "stringconstant":
		var term StringConstant
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
		// multiset operators
	case "add":
		var term MultisetAdd
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "all":
		var term MultisetAll
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "numberof":
		var term MultisetNumberOf
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "subtract":
		var term MultisetSubtract
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "scalarproduct":
		var term MultisetScalarProduct
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "empty":
		var term MultisetEmpty
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
		// Other operators
	case "tuple":
		var term HLTupleOperator
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	case "useroperator":
		var term HLUserOperator
		if err := d.DecodeElement(&term, &start); err != nil {
			return err
		}
		h.Value = term
	default:
		log.Panic("HLTerm: Unknown term ", h.Type)
	}
	return nil
}

type HLSubterm struct {
	XMLName xml.Name `xml:"subterm"`
	Term    HLTerm   `xml:",any"`
}

func (h HLSubterm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpH struct {
		XMLName xml.Name `xml:"subterm"`
		Term    interface{}
	}

	t := tmpH{h.XMLName, h.Term.Value}

	return e.Encode(t)
}

// Booleans

type BoolEquality struct {
	XMLName xml.Name    `xml:"equality"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolInequality struct {
	XMLName xml.Name    `xml:"inequality"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolAnd struct {
	XMLName xml.Name    `xml:"and"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolOr struct {
	XMLName xml.Name    `xml:"or"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolImply struct {
	XMLName xml.Name    `xml:"imply"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolNot struct {
	XMLName xml.Name    `xml:"not"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolConstant struct {
	XMLName xml.Name    `xml:"booleanconstant"`
	Value   *bool       `xml:"value,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type BoolSort struct {
	XMLName xml.Name `xml:"bool"`
}

// Finite enumeration

type FEConstant struct {
	XMLName xml.Name `xml:"feconstant"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
}

type FESort struct {
	XMLName   xml.Name     `xml:"finiteenumeration"`
	Constants []FEConstant `xml:"feconstant"` // optional
}

// Cyclic enumeration

type CyclicEnumSort struct {
	XMLName   xml.Name     `xml:"cyclicenumeration"`
	Constants []FEConstant `xml:"feconstant"` // optional
}

type CyclicEnumSuccessor struct {
	XMLName xml.Name    `xml:"successor"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type CyclicEnumPredecessor struct {
	XMLName xml.Name    `xml:"predecessor"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

// Finite integer ranges

type FIRSort struct {
	XMLName xml.Name `xml:"finiteintrange"`
	Start   *int     `xml:"start,attr"`
	End     *int     `xml:"end,attr"`
}

type FIRLessThan struct {
	XMLName xml.Name    `xml:"lessthan"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type FIRLessThanOrEqual struct {
	XMLName xml.Name    `xml:"lessthanorequal"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type FIRGreaterThan struct {
	XMLName xml.Name    `xml:"greaterthan"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type FIRGreaterThanOrEqual struct {
	XMLName xml.Name    `xml:"greaterthanorequal"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type FIRConstant struct {
	XMLName xml.Name    `xml:"finiteintrangeconstant"`
	Value   *int        `xml:"value,attr"`
	FIRSort *FIRSort    `xml:"finiteintrange"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

// Dots

type DotSort struct {
	XMLName xml.Name `xml:"dot"`
}

type DotConstant struct {
	XMLName xml.Name `xml:"dotconstant"`
}

// Multisets

type MultisetCardinality struct {
	XMLName xml.Name    `xml:"cardinality"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetCardinalityOf struct {
	XMLName xml.Name    `xml:"cardinalityof"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetContains struct {
	XMLName xml.Name    `xml:"contains"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetAdd struct {
	XMLName xml.Name    `xml:"add"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetSubtract struct {
	XMLName xml.Name    `xml:"subtract"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetAll struct {
	XMLName xml.Name    `xml:"all"`
	Terms   []HLSubterm `xml:"subterm"` // optional
	Sort    *HLSort     `xml:",any"`
}

func (m MultisetAll) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpM struct {
		XMLName xml.Name    `xml:"all"`
		Terms   []HLSubterm `xml:"subterm"`
		Sort    interface{}
	}

	t := tmpM{m.XMLName, m.Terms, m.Sort.Value}

	return e.Encode(t)
}

type MultisetEmpty struct {
	XMLName xml.Name    `xml:"empty"`
	Terms   []HLSubterm `xml:"subterm"` // optional
	Sort    *HLSort     `xml:",any"`
}

func (m MultisetEmpty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpM struct {
		XMLName xml.Name    `xml:"empty"`
		Terms   []HLSubterm `xml:"subterm"`
		Sort    interface{}
	}

	t := tmpM{m.XMLName, m.Terms, m.Sort.Value}

	return e.Encode(t)
}

type MultisetScalarProduct struct {
	XMLName xml.Name    `xml:"scalarproduct"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type MultisetNumberOf struct {
	XMLName xml.Name    `xml:"numberof"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

// Partitions

type PartitionSortDeclaration struct {
	XMLName xml.Name `xml:"partition"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	HLSort
	PartitionElements []PartitionElement `xml:"partitionelement"` // one or more
}

type PartitionElement struct {
	XMLName xml.Name `xml:"partitionelement"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	Terms   []HLTerm `xml:",any"`
}

type PartitionLessThan struct {
	XMLName xml.Name    `xml:"ltp"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type PartitionGreaterThan struct {
	XMLName xml.Name    `xml:"gtp"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type PartitionElementOf struct {
	XMLName xml.Name    `xml:"partitionelementof"`
	Ref     *string     `xml:"refpartition,attr"` // data of type IDREF
	Terms   []HLSubterm `xml:"subterm"`           // optional
}

// Lists

type ListSort struct {
	XMLName xml.Name `xml:"list"`
	Sort    *HLSort  `xml:",any"`
}

func (l ListSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"list"`
		Sort    interface{}
	}

	t := tmpL{l.XMLName, l.Sort.Value}

	return e.Encode(t)
}

type ListAppend struct {
	XMLName xml.Name    `xml:"listappend"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type ListConcatenation struct {
	XMLName xml.Name    `xml:"listconcatenation"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type ListMake struct {
	XMLName xml.Name    `xml:"makelist"`
	Sort    *HLSort     `xml:",any"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (l ListMake) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type tmpL struct {
		XMLName xml.Name `xml:"makelist"`
		Sort    interface{}
		Terms   []HLSubterm `xml:"subterm"`
	}

	t := tmpL{l.XMLName, l.Sort.Value, l.Terms}

	return e.Encode(t)
}

type ListLength struct {
	XMLName xml.Name    `xml:"listlength"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type ListMemberAtIndex struct {
	XMLName xml.Name    `xml:"memberatindex"`
	Index   *uint       `xml:"index,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

type ListSublist struct {
	XMLName xml.Name    `xml:"sublist"`
	Start   *uint       `xml:"start,attr"`
	Length  *uint       `xml:"length,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

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
