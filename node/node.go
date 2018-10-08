package node

import (
	"github.com/0shima_x/go-cart/splitter"
)

type Node struct {
	Left      interface{}
	Right     interface{}
	Col       int
	Predicate splitter.Predicate
	Value     interface{}
	Results   map[string]int
}

// const (
// 	NUM_ND = "NUMERIC"
// 	CAT_ND = "CATEGORICAl"
// )
//
// type Node interface {
// 	Type() string // ObjectType
// 	Inspect() string
// }
//
// type NumNode struct {
// 	Col       int32
// 	Value     float64
// 	Results   bool
// 	Predicate bool
// }
//
// func (nn *NumNode) Type() string    { return NUM_ND }
// func (nn *NumNode) Inspect() string { return strconv.FormatFloat(nn.Value, 'f', 4, 64) }
//
// type CatNode struct {
// 	Col       int32
// 	Value     string
// 	Results   bool
// 	Predicate bool
// }
//
// func (cn *CatNode) Type() string    { return CAT_ND }
// func (cn *CatNode) Inspect() string { return cn.Value }

// node1 Node := &NumNode{}
// node2 Node := &CatNode{}
