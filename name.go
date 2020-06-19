package main

import (
	"encoding/xml"
)

type Name struct {
	XMLName       xml.Name            `xml:"name"`
	Text          *string             `xml:"text"`         // optional
	Graphics      *AnnotationGraphics `xml:"graphics"`     // optional
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"` // optional
}
