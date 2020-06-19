package main

import "encoding/xml"

type Pnml struct {
	XMLName xml.Name `xml:"pnml"`
	XMLNs   string   `xml:"xmlns,attr"`
	Nets    []Net    `xml:"net"` // at least 1
}

type Net struct {
	XMLName       xml.Name       `xml:"net"`
	ID            string         `xml:"id,attr"`
	Type          string         `xml:"type,attr"`
	Name          *Name          `xml:"name"`
	Pages         []Page         `xml:"page"` // at least 1
	ToolSpecifics []ToolSpecific `xml:"toolspecific"`
	// net.labels def to be empty
}

type Name struct {
	XMLName       xml.Name            `xml:"name"`
	Text          *string             `xml:"text"`
	Graphics      *AnnotationGraphics `xml:"graphics"`
	ToolSpecifics []ToolSpecific      `xml:"toolspecific"`
}

type Page struct {
	XMLName        xml.Name        `xml:"page"`
	ID             string          `xml:"id,attr"`
	Name           *Name           `xml:"name"`
	ToolSpecifics  []ToolSpecific  `xml:"toolspecific"`
	Pages          []Page          `xml:"page"`
	Places         []Place         `xml:"place"`
	Transitions    []Transition    `xml:"transition"`
	RefPlaces      []RefPlace      `xml:"referencePlace"`
	RefTransitions []RefTransition `xml:"referenceTransition"`
	Arcs           []Arc           `xml:"arc"`
	Graphics       *PageGraphics   `xml:"graphics"`
	// page.labels def to be empty
}

type ToolSpecific struct {
	XMLName xml.Name `xml:"toolspecific"`
	Tool    string   `xml:"tool,attr"`
	Version string   `xml:"version,attr"`
	// Zero or more anyElement
}

type Place struct {
	XMLName        xml.Name       `xml:"place"`
	ID             string         `xml:"id,attr"`
	Name           *Name          `xml:"name"`
	ToolSpecifics  []ToolSpecific `xml:"toolspecific"`
	InitialMarking *PTMarking     `xml:"initialMarking"` // label for pt nets only
	Graphics       *NodeGraphics  `xml:"graphics"`
	// place.labels defined to be empty
}

type Transition struct {
	XMLName       xml.Name       `xml:"transition"`
	ID            string         `xml:"id,attr"`
	Name          *Name          `xml:"name"`
	ToolSpecifics []ToolSpecific `xml:"toolspecific"`
	Graphics      *NodeGraphics  `xml:"graphics"`
	// transition.labels def to be empty
}

type RefPlace struct {
	XMLName       xml.Name       `xml:"referencePlace"`
	ID            string         `xml:"id,attr"`
	Reference     string         `xml:"ref,attr"`
	Name          *Name          `xml:"name"`
	ToolSpecifics []ToolSpecific `xml:"toolspecific"`
	Graphics      *NodeGraphics  `xml:"graphics"`
}

type RefTransition struct {
	XMLName       xml.Name       `xml:"referenceTransition"`
	ID            string         `xml:"id,attr"`
	Reference     string         `xml:"ref,attr"`
	Name          *Name          `xml:"name"`
	ToolSpecifics []ToolSpecific `xml:"toolspecific"`
	Graphics      *NodeGraphics  `xml:"graphics"`
}

type Arc struct {
	XMLName       xml.Name         `xml:"arc"`
	ID            string           `xml:"id,attr"`
	Source        string           `xml:"source,attr"`
	Target        string           `xml:"target,attr"`
	Name          *Name            `xml:"name"`
	Weight        *PTArcAnnotation `xml:"inscription"` // label for pt nets only
	Graphics      *EdgeGraphics    `xml:"graphics"`
	ToolSpecifics []ToolSpecific   `xml:"toolspecific"`
	// arc.labels defined to be empty
}

type PageGraphics = NodeGraphics

type NodeGraphics struct {
	XMLName   xml.Name   `xml:"graphics"`
	Position  Position   `xml:"position"`
	Dimension *Dimension `xml:"dimension"`
	Fill      *Fill      `xml:"fill"`
	Line      *Line      `xml:"line"`
}

type EdgeGraphics struct {
	XMLName   xml.Name   `xml:"graphics"`
	Positions []Position `xml:"position"`
	Line      *Line      `xml:"line"`
}

type Position struct {
	XMLName xml.Name `xml:"position"`
	X       float32  `xml:"x,attr"`
	Y       float32  `xml:"y,attr"`
}

type Dimension struct {
	XMLName xml.Name `xml:"dimension"`
	X       float32  `xml:"x,attr"` // between 0 and 999.9, 4 digits with 1 decimal digit
	Y       float32  `xml:"y,attr"` // between 0 and 999.9, 4 digits with 1 decimal digit
}

type Fill struct {
	XMLName          xml.Name `xml:"fill"`
	Color            *string  `xml:"color,attr"`
	GradientColor    *string  `xml:"gradient-color,attr"`
	GradientRotation *string  `xml:"gradient-rotation,attr"` // must be vertical, horizontal, or diagonal
	Image            *string  `xml:"image,attr"`
}

type Line struct {
	XMLName xml.Name `xml:"line"`
	Shape   *string  `xml:"shape,attr"` // must be line or curve
	Color   *string  `xml:"color,attr"`
	Width   *float32 `xml:"width,attr"` // between 0 and 999.9, 4 digits with 1 decimal digit
	Style   *string  `xml:"style,attr"` // must be solid, dash, or dot
}

type AnnotationGraphics struct {
	XMLName xml.Name `xml:"graphics"`
	Offset  Offset   `xml:"offset"`
	Fill    *Fill    `xml:"fill"`
	Line    *Line    `xml:"line"`
	Font    *Font    `xml:"font"`
}

type Offset struct {
	XMLName xml.Name `xml:"offset"`
	X       float32  `xml:"x,attr"`
	Y       float32  `xml:"y,attr"`
}

type Font struct {
	XMLName    xml.Name `xml:"font"`
	Family     *string  `xml:"family,attr"`
	Style      *string  `xml:"style,attr"`
	Weight     *string  `xml:"weight,attr"`
	Size       *string  `xml:"size,attr"`
	Decoration *string  `xml:"decoration,attr"` // must be underline, overline, or line-through
	Align      *string  `xml:"align,attr"`      // must be left, center, or right
	Rotation   *float32 `xml:"rotation,attr"`
}
