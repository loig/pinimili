package pnml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

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
		line, col := d.InputPos()
		msg := fmt.Sprint(modelPath, " at line ", line, ", col ", col, ", unknown term ", h.Type)
		if panicIfNotPnmlCompliant {
			return errors.New(msg)
		}
		switch h.Type {
		case "subterm":
			log.Print("Pinimili: ", msg)
			var term HLSubterm
			if err := d.DecodeElement(&term, &start); err != nil {
				return err
			}
			h.Value = term
		default:
			return errors.New(msg)
		}
		/*
			var term HLSubterm
			if err := d.DecodeElement(&term, &start); err != nil {
				return err
			}
			h.Value = term.Term.Value
		*/
	}
	return nil
}
