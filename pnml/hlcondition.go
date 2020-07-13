package pnml

import "encoding/xml"

type HLCondition struct {
	XMLName       xml.Name            `xml:"condition"`
	Info          *string             `xml:"text"`         // optional
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}
