package node

type Link int

const (
	LEAF Link = iota
	BRANCH
)

func (l Link) String() string {
	switch l {
	case LEAF:
		return "Leaf"
	case BRANCH:
		return "Branch"
	}
	panic("encounted a unkonwn node type.")
}

type Node struct {
	Left       interface{}
	Right      interface{}
	SplitCol   int
	SplitValue interface{}
	Results    map[string]int
	NodeType   Link
}
