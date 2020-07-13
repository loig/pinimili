package pnml

import "encoding/xml"

type HLTupleOperator struct {
	XMLName xml.Name    `xml:"tuple"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}
