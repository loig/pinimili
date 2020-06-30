package pnml

import "encoding/xml"

type HLDeclaration struct {
	XMLName              xml.Name                `xml:"declaration"`
	Info                 *string                 `xml:"text"`                                 // optional
	SortDeclarations     []HLSortDeclaration     `xml:"structure>declarations>namedsort"`     // optional
	VariableDeclarations []HLVariableDeclaration `xml:"structure>declarations>variabledecl"`  // optional
	OperatorDeclarations []HLOperatorDeclaration `xml:"structure>declarations>namedoperator"` // optional
}

type HLVariableDeclaration struct {
	XMLName xml.Name `xml:"variabledecl"`
	ID      *string  `xml:"id,attr"`
	Name    *string  `xml:"name,attr"`
	// start choice
	BoolSort *BoolSort `xml:"bool"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type HLMultisetSort struct {
	XMLName xml.Name `xml:"multisetsort"`
	// start choice
	BoolSort *BoolSort `xml:"bool"`
	//<ref name="BuiltInSort"/>
	MultisetSort *HLMultisetSort `xml:"multisetsort"`
	ProductSort  *HLProductSort  `xml:"productsort"`
	UserSort     *HLUserSort     `xml:"usersort"`
	// end choice
}

type HLProductSort struct {
	XMLName  xml.Name   `xml:"productsort"`
	BoolSort []BoolSort `xml:"bool"` // optional
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
	BoolSort *BoolSort `xml:"bool"`
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
	BoolSort *BoolSort `xml:"structure>bool"` // optional
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
	Variable   *HLVariable     `xml:"variable"`
	Equality   *BoolEquality   `xml:"equality"`
	Inequality *BoolInequality `xml:"inequality"`
	And        *BoolAnd        `xml:"and"`
	Or         *BoolOr         `xml:"or"`
	Imply      *BoolImply      `xml:"imply"`
	Not        *BoolNot        `xml:"not"`
	//<ref name="BuiltInOperator"/>
	Boolean *BoolConstant `xml:"booleanconstant"`
	//<ref name="BuiltInConstant"/>
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
}

type BoolSort struct {
	XMLName xml.Name `xml:"bool"`
}
