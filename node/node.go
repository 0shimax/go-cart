package node

import "fmt"

const (
	LEAF   = "LEAF"
	BRANCH = "BRANCH"
)

type Node interface {
	String() string
	Print() string
	GetLeft() Node
	GetRight() Node
}

type Leaf struct {
	Results map[string]int
}

func (l Leaf) String() string { return LEAF }
func (l Leaf) Print() string  { return fmt.Sprint(l.Results) }
func (l Leaf) GetLeft() Node  { return l }
func (l Leaf) GetRight() Node { return l }

type Branch struct {
	Left       Node
	Right      Node
	SplitCol   int
	SplitValue interface{}
}

func (b Branch) String() string { return BRANCH }
func (b Branch) Print() string  { return fmt.Sprint(b.SplitCol) + ":" + fmt.Sprint(b.SplitValue) + "?" }
func (b Branch) GetLeft() Node  { return b.Left }
func (b Branch) GetRight() Node { return b.Right }
