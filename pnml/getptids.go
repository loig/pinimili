package pnml

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"

	"golang.org/x/net/html/charset"
)

// Getptids returns the ids of the places and transitions of a pnml model
// described in the file whose path is given in argument, if the panicOnWrongModel
// argument is false, the model will be read even if it contains small mistakes
// such as wrong operators arities
func Getptids(path string, panicOnWrongModel bool) (pids []string, tids []string) {

	panicIfNotPnmlCompliant = panicOnWrongModel

	pnmlFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(pnmlFile)
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

	pids, tids = pnml.getptids()

	return pids, tids
}

func (p Pnml) getptids() (pids []string, tids []string) {

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
