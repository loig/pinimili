package pnml

import "encoding/xml"

type HLDeclaration struct {
	XMLName                xml.Name                `xml:"declaration"`
	Info                   *string                 `xml:"text"`                                 // optional
	SortDeclarations       []HLSortDeclaration     `xml:"structure>declarations>namedsort"`     // optional
	VariableDeclarations   []HLVariableDeclaration `xml:"structure>declarations>variabledecl"`  // optional
	OperatorDeclarations   []HLOperatorDeclaration `xml:"structure>declarations>namedoperator"` // optional
	FEConstantDeclarations []FEConstant            `xml:"structure>declarations>feconstant"`    // optional
}

type HLVariableDeclaration struct {
	XMLName xml.Name `xml:"variabledecl"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	// start choice
	BoolSort       *BoolSort       `xml:"bool"`
	FESort         *FESort         `xml:"finiteenumeration"`
	CyclicEnumSort *CyclicEnumSort `xml:"cyclicenumeration"`
	FIRSort        *FIRSort        `xml:"finiteintrange"`
	DotSort        *DotSort        `xml:"dot"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type HLMultisetSort struct {
	XMLName xml.Name `xml:"multisetsort"`
	// start choice
	BoolSort       *BoolSort       `xml:"bool"`
	FESort         *FESort         `xml:"finiteenumeration"`
	CyclicEnumSort *CyclicEnumSort `xml:"cyclicenumeration"`
	FIRSort        *FIRSort        `xml:"finiteintrange"`
	DotSort        *DotSort        `xml:"dot"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type HLProductSort struct {
	XMLName        xml.Name         `xml:"productsort"`
	BoolSort       []BoolSort       `xml:"bool"`              // optional
	FESort         []FESort         `xml:"finiteenumeration"` // optional
	CyclicEnumSort []CyclicEnumSort `xml:"cyclicenumeration"` // optional
	FIRSort        []FIRSort        `xml:"finiteintrange"`    // optional
	DotSort        []DotSort        `xml:"dot"`               // optional
	//<ref name="BuiltInSort"/>  0 or more
	MultisetSort []HLMultisetSort `xml:"multisetsort"` // optional
	ProductSort  []HLProductSort  `xml:"productsort"`  // optional
	UserSort     []HLUserSort     `xml:"usersort"`     // optional
}

type HLUserSort struct {
	XMLName xml.Name `xml:"usersort"`
	ID      *string  `xml:"declaration,attr"` // data type="IDREF"
}

type HLSortDeclaration struct {
	XMLName xml.Name `xml:"namedsort"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	// start choice
	BoolSort       *BoolSort       `xml:"bool"`
	FESort         *FESort         `xml:"finiteenumeration"`
	CyclicEnumSort *CyclicEnumSort `xml:"cyclicenumeration"`
	FIRSort        *FIRSort        `xml:"finiteintrange"`
	DotSort        *DotSort        `xml:"dot"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type HLOperatorDeclaration struct {
	XMLName              xml.Name                `xml:"namedoperator"`
	ID                   *string                 `xml:"id,attr"`
	Name                 *string                 `xml:"name,attr"`
	VariableDeclarations []HLVariableDeclaration `xml:"parameter>variabledecl"` // optional
	Def                  *HLTerm                 `xml:"def"`
}

type HLVariable struct {
	XMLName xml.Name `xml:"variable"`
	ID      *string  `xml:"refvariable,attr"` // data type="IDREF"
}

type HLTupleOperator struct {
	XMLName xml.Name `xml:"tuple"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type HLUserOperator struct { // recursion forbidden
	XMLName xml.Name `xml:"useroperator"`
	ID      *string  `xml:"declaration,attr"` // data type="IDREF"
	Terms   []HLTerm `xml:"subterm"`          // optional
}

type HLType struct {
	XMLName       xml.Name            `xml:"type"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	// start choice
	BoolSort       *BoolSort       `xml:"structure>bool"`              // optional
	FESort         *FESort         `xml:"structure>finiteenumeration"` // optional
	CyclicEnumSort *CyclicEnumSort `xml:"structure>cyclicenumeration"` // optional
	FIRSort        *FIRSort        `xml:"structure>finiteintrange"`    // optional
	DotSort        *DotSort        `xml:"structure>dot"`               // optional
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"structure>multisetsort"` // optional
	ProductSort  *HLProductSort  `xml:"structure>productsort"`  // optional
	UserSort     *HLUserSort     `xml:"structure>usersort"`     // optional
	// end choice
}

type HLMarking struct {
	XMLName       xml.Name            `xml:"hlinitialMarking"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Term          *HLTerm             `xml:"structure"`    // optional
}

type HLCondition struct {
	XMLName       xml.Name            `xml:"condition"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Term          *HLTerm             `xml:"structure"`    // optional
}

type HLAnnotation struct {
	XMLName       xml.Name            `xml:"hlinscription"`
	Info          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Term          *HLTerm             `xml:"structure"`    // optional
}

type HLTerm struct {
	// choice
	Variable              *HLVariable            `xml:"variable"`
	Equality              *BoolEquality          `xml:"equality"`
	Inequality            *BoolInequality        `xml:"inequality"`
	And                   *BoolAnd               `xml:"and"`
	Or                    *BoolOr                `xml:"or"`
	Imply                 *BoolImply             `xml:"imply"`
	Not                   *BoolNot               `xml:"not"`
	Predecessor           *CyclicEnumPredecessor `xml:"predecessor"`
	Successor             *CyclicEnumSuccessor   `xml:"successor"`
	FIRLessThan           *FIRLessThan           `xml:"lessthan"`
	FIRLessThanOrEqual    *FIRLessThanOrEqual    `xml:"lessthanorequal"`
	FIRGreaterThan        *FIRGreaterThan        `xml:"greaterthan"`
	FIRGreaterThanOrEqual *FIRGreaterThanOrEqual `xml:"greaterthanorequal"`
	MultisetCardinality   *MultisetCardinality   `xml:"cardinality"`
	MultisetCardinalityOf *MultisetCardinalityOf `xml:"cardinalityof"`
	MultisetContains      *MultisetContains      `xml:"contains"`
	//<ref name="BuiltInOperator"/>
	BoolConstant *BoolConstant `xml:"booleanconstant"`
	FIRConstant  *FIRConstant  `xml:"finiteintrangeconstant"`
	DotConstant  *DotConstant  `xml:"dotconstant"`
	//<ref name="BuiltInConstant"/>
	MultisetAdd           *MultisetAdd           `xml:"add"`
	MultisetAll           *MultisetAll           `xml:"all"`
	MultisetNumberOf      *MultisetNumberOf      `xml:"numberof"`
	MultisetSubtract      *MultisetSubtract      `xml:"subtract"`
	MultisetScalarProduct *MultisetScalarProduct `xml:"scalarproduct"`
	MultisetEmpty         *MultisetEmpty         `xml:"empty"`
	//<ref name="MultisetOperator"/>
	TupleOperator *HLTupleOperator `xml:"tuple"`
	UserOperator  *HLUserOperator  `xml:"useroperator"`
}

// Booleans

type BoolEquality struct {
	XMLName xml.Name `xml:"equality"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolInequality struct {
	XMLName xml.Name `xml:"inequality"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolAnd struct {
	XMLName xml.Name `xml:"and"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolOr struct {
	XMLName xml.Name `xml:"or"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolImply struct {
	XMLName xml.Name `xml:"imply"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolNot struct {
	XMLName xml.Name `xml:"not"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type BoolConstant struct {
	XMLName xml.Name `xml:"booleanconstant"`
	Value   bool     `xml:"value,attr"`
	Terms   []HLTerm `xml:"subterm"` // optional
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
	XMLName xml.Name `xml:"successor"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type CyclicEnumPredecessor struct {
	XMLName xml.Name `xml:"predecessor"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

// Finite integer ranges

type FIRSort struct {
	XMLName xml.Name `xml:"finiteintrange"`
	Start   *int     `xml:"start,attr"`
	End     *int     `xml:"end,attr"`
}

type FIRLessThan struct {
	XMLName xml.Name `xml:"lessthan"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type FIRLessThanOrEqual struct {
	XMLName xml.Name `xml:"lessthanorequal"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type FIRGreaterThan struct {
	XMLName xml.Name `xml:"greaterthan"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type FIRGreaterThanOrEqual struct {
	XMLName xml.Name `xml:"greaterthanorequal"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type FIRConstant struct {
	XMLName xml.Name `xml:"finiteintrangeconstant"`
	Value   *int     `xml:"value,attr"`
	FIRSort *FIRSort `xml:"finiteintrange"`
	Terms   []HLTerm `xml:"subterm"` // optional
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
	XMLName xml.Name `xml:"cardinality"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetCardinalityOf struct {
	XMLName xml.Name `xml:"cardinalityof"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetContains struct {
	XMLName xml.Name `xml:"contains"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetAdd struct {
	XMLName xml.Name `xml:"add"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetSubtract struct {
	XMLName xml.Name `xml:"subtract"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetAll struct {
	XMLName xml.Name `xml:"all"`
	Terms   []HLTerm `xml:"subterm"` // optional
	// start choice
	BoolSort       *BoolSort       `xml:"bool"`
	FESort         *FESort         `xml:"finiteenumeration"`
	CyclicEnumSort *CyclicEnumSort `xml:"cyclicenumeration"`
	FIRSort        *FIRSort        `xml:"finiteintrange"`
	DotSort        *DotSort        `xml:"dot"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type MultisetEmpty struct {
	XMLName xml.Name `xml:"empty"`
	Terms   []HLTerm `xml:"subterm"` // optional
	// start choice
	BoolSort       *BoolSort       `xml:"bool"`
	FESort         *FESort         `xml:"finiteenumeration"`
	CyclicEnumSort *CyclicEnumSort `xml:"cyclicenumeration"`
	FIRSort        *FIRSort        `xml:"finiteintrange"`
	DotSort        *DotSort        `xml:"dot"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type MultisetScalarProduct struct {
	XMLName xml.Name `xml:"scalarproduct"`
	Terms   []HLTerm `xml:"subterm"` // optional
}

type MultisetNumberOf struct {
	XMLName xml.Name `xml:"numberof"`
	Terms   []HLTerm `xml:"subterm"` // optional
}