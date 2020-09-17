package pnml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Getptids returns the ids of the places and transitions of a pnml model
// described in the file whose path is given in argument
func Getptids(path string) (pids []string, tids []string) {

	pnmlFile, err := os.Open("test.pnml")
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(pnmlFile)
	if err != nil {
		panic(err)
	}

	var pnml Pnml
	err = xml.Unmarshal(byteValue, &pnml)
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
