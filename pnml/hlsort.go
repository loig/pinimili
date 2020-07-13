package pnml

import (
	"encoding/xml"
	"log"
)

type HLSort struct {
	Type  string `xml:"-"`
	Value interface{}
}

func (h *HLSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	h.Type = start.Name.Local
	switch h.Type {
	// built in sorts
	case "bool":
		var sort BoolSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "finiteenumeration":
		var sort FESort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "cyclicenumeration":
		var sort CyclicEnumSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "finiteintrange":
		var sort FIRSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "dot":
		var sort DotSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "list":
		var sort ListSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "natural":
		var sort IntNatural
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "positive":
		var sort IntPositive
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "integer":
		var sort IntInteger
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	case "string":
		var sort StringSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
		// multiset sort
	case "multisetsort":
		var sort HLMultisetSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
		// product sort
	case "productsort":
		var sort HLProductSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
		// user defined sort
	case "usersort":
		var sort HLUserSort
		if err := d.DecodeElement(&sort, &start); err != nil {
			return err
		}
		h.Value = sort
	default:
		log.Panic("HLSort: Unknown sort ", h.Type)
	}
	return nil
}
