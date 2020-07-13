package pnml

import (
	"encoding/xml"
)

type HLType struct {
	XMLName       xml.Name            `xml:"type"`
	Info          *string             `xml:"text"`         // optional
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
	Structure     *HLSortStructure    `xml:"structure"`    // optional
}
