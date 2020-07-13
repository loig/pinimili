package pnml

import (
	"encoding/xml"
)

type HLProductSort struct {
	XMLName xml.Name `xml:"productsort"`
	Sorts   []HLSort `xml:",any"` // optional
}

func (h HLProductSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	startToken := xml.StartElement{Name: xml.Name{Space: "", Local: "productsort"}}

	err := e.EncodeToken(startToken)
	if err != nil {
		return err
	}

	for _, sort := range h.Sorts {
		e.Encode(sort.Value)
	}

	return e.EncodeToken(startToken.End())
}
