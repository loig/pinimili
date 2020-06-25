package ptnet

type Place struct {
	ID      string
	Name    string
	Marking uint64
}

type Transition struct {
	ID   string
	Name string
}

type Arc struct {
	ID       string
	Weight   uint64
	FromID   string
	FromType NodeType
	ToID     string
	ToType   NodeType
}

type NodeType int

const (
	PlaceNode NodeType = iota
	TransitionNode
)

type Net struct {
	Places      map[string]Place
	Transitions map[string]Transition
	Arcs        []Arc
}

func New() *Net {
	var n Net
	n.Places = make(map[string]Place)
	n.Transitions = make(map[string]Transition)
	n.Arcs = make([]Arc, 0)
	return &n
}
