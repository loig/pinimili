package pnml

import "encoding/xml"

type HLAnnotation struct {
	XMLName       xml.Name            `xml:"hlinscription"`
	Info          *string             `xml:"text"`         // optional
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLTermStructure    `xml:"structure"`    // optional
}
