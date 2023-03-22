package pnml

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

// GetPnml parses a pnml model described in the file whose path is given in
// argument, if the panicOnWrongModel argument is false, the model will be read
// even if it contains small mistakes such as wrong operators arities
func GetPnml(path string, panicOnWrongModel bool) *Pnml {

	panicIfNotPnmlCompliant = panicOnWrongModel
	modelPath = path

	pnmlFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	byteValue, err := io.ReadAll(pnmlFile)
	if err != nil {
		panic(err)
	}

	var pnml Pnml

	reader := bytes.NewReader(byteValue)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&pnml)

	// err = xml.Unmarshal(byteValue, &pnml)

	if err != nil {
		panic(err)
	}

	return &pnml
}

// Getptids returns the ids of the places and transitions of a pnml model
func (p Pnml) Getptids() (pids []string, tids []string) {

	pids = make([]string, 0)
	tids = make([]string, 0)

	for _, n := range p.Nets {
		for _, pa := range n.Pages {
			for _, p := range pa.Places {
				pids = append(pids, *(p.ID))
			}
			for _, t := range pa.Transitions {
				tids = append(tids, *(t.ID))
			}
		}
	}

	return pids, tids
}

// GetMaxConstantInMarking gives the maximum constant appearing in the initial
// marking of a net represented by a pnml model (only for PT nets)
func (p Pnml) GetMaxConstantInMarking() (c uint64) {
	for _, n := range p.Nets {
		for _, pa := range n.Pages {
			for _, p := range pa.Places {
				if p.InitialMarking != nil {
					if p.InitialMarking.Tokens != nil {
						if *(p.InitialMarking.Tokens) > c {
							c = *(p.InitialMarking.Tokens)
						}
					}
				}
			}
		}
	}
	return c
}
