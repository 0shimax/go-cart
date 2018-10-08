package main

import (
	"github.com/0shima_x/go-cart/data_loader"
	"github.com/0shima_x/go-cart/tree"
)

func main() {
	path := "./data.csv"
	datas := data_loader.Load(path)
	// fmt.Printf("Data: %v\n", datas)

	X, y := data_loader.SplitIntoFeaturesLabes(datas, 4)
	// fmt.Printf("X: %v\n", X)
	// fmt.Printf("y: %v\n", y)

	// setX1, setX2, setY1, setY2 := splitter.Split(X, y, 2, "yes")
	// fmt.Printf("setX1: %v\n", setX1)
	// fmt.Printf("setX2: %v\n", setX2)
	// fmt.Printf("setY1: %v\n", setY1)
	// fmt.Printf("setY2: %v\n", setY2)
	//
	// gini := criterion.GiniImpurity(setY1)
	// ent := criterion.Entropy(setY1)
	// fmt.Printf("gini: %f\n", gini)
	// fmt.Printf("ent: %f\n", ent)

	model := tree.DecisionTreeClassifier{Criterion: "entropy"}
	tree := model.Build(X, y)
	model.Print(tree, " ")
}
