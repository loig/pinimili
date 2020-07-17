package pnml

import (
	"encoding/xml"
	"errors"
)

type StringSubstring struct {
	XMLName xml.Name    `xml:"substring"`
	Start   *uint       `xml:"start,attr"`
	Length  *uint       `xml:"length,attr"`
	Terms   []HLSubterm `xml:"subterm"` // optional
}

func (s *StringSubstring) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type stringsubstring StringSubstring
	var ss stringsubstring
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	if ss.Start == nil {
		return errors.New("StringSubstring: a substring must have a start attribute")
	}
	if ss.Length == nil {
		return errors.New("StringSubstring: a substring must have a length attribute")
	}
	if len(ss.Terms) != 1 {
		return errors.New("StringSubstring: a substring must have exactly one subterm")
	}
	*s = StringSubstring(ss)
	return nil
}
