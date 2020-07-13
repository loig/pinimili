package pnml

import "encoding/xml"

type HLMarking struct {
	XMLName       xml.Name            `xml:"hlinitialMarking"`
	Info          *string             `xml:"text"`         // optional
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}
