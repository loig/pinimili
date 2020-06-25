package ptnet

import (
	"log"

	"github.com/loig/pinimili/pnml"
)

func NewFromPnml(p pnml.Pnml) []Net {
	// TODO:
	// - places/transitions with same ID
	// - get all pages (pages in pages)

	log.SetPrefix("ptnet.NewFromPnml:")

	ptNets := make([]Net, 0)
	for _, net := range p.Nets {
		if *net.Type != "http://www.pnml.org/version-2009/grammar/ptnet" {
			log.Print("Ignored one net (wrong type)")
			continue
		}

		var ptNet Net

		// Get the places
		ptNet.Places = make(map[string]Place)
		for _, page := range net.Pages {
			for _, place := range page.Places {
				var ptPlace Place
				ptPlace.ID = *place.ID
				if place.Name != nil && (*place.Name).Text != nil {
					ptPlace.Name = *(*place.Name).Text
				}
				if place.InitialMarking != nil {
					ptPlace.Marking = *(*place.InitialMarking).Tokens
				}
				ptNet.Places[ptPlace.ID] = ptPlace
			}
		}

		// Get the transitions
		ptNet.Transitions = make(map[string]Transition)
		for _, page := range net.Pages {
			for _, transition := range page.Transitions {
				var ptTransition Transition
				ptTransition.ID = *transition.ID
				if transition.Name != nil && (*transition.Name).Text != nil {
					ptTransition.Name = *(*transition.Name).Text
				}
				ptNet.Transitions[ptTransition.ID] = ptTransition
			}
		}

		// Get the arcs
		ptNet.Arcs = make([]Arc, 0)
		for _, page := range net.Pages {
			for _, arc := range page.Arcs {
				var ptArc Arc
				ptArc.ID = *arc.ID
				if arc.Weight == nil {
					log.Panic("An arc in a PTNet must have a weight")
				}
				ptArc.Weight = *(*arc.Weight).Value

				// Arc source
				ptArc.FromID = *arc.Source
				_, isPlace := ptNet.Places[ptArc.FromID]
				if isPlace {
					ptArc.FromType = PlaceNode
				} else {
					_, isTransition := ptNet.Transitions[ptArc.FromID]
					if !isTransition {
						log.Panic("Arc ", ptArc.ID, " has unknown source ", ptArc.FromID)
					}
					ptArc.FromType = TransitionNode
				}

				// Arc target
				ptArc.ToID = *arc.Target
				_, isPlace = ptNet.Places[ptArc.ToID]
				if isPlace {
					ptArc.ToType = PlaceNode
				} else {
					_, isTransition := ptNet.Transitions[ptArc.ToID]
					if !isTransition {
						log.Panic("Arc ", ptArc.ID, " has unknown target ", ptArc.ToID)
					}
					ptArc.ToType = TransitionNode
				}

				ptNet.Arcs = append(ptNet.Arcs, ptArc)
			}
		}

		ptNets = append(ptNets, ptNet)
	}

	return ptNets
}
