package main

import "encoding/xml"

type PTMarking struct {
	XMLName       xml.Name            `xml:"initialMarking"`
	Tokens        uint64              `xml:"text"` // must be integer >= 0
	Graphics      *AnnotationGraphics `xml:"graphics"`
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"`
}

type PTArcAnnotation struct {
	XMLName       xml.Name            `xml:"inscription"`
	Value         uint64              `xml:"text"` // must be integer > 0
	Graphics      *AnnotationGraphics `xml:"graphics"`
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"`
}
