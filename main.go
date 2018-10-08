package main

import (
	"fmt"

	"github.com/0shima_x/go-cart/criterion"
	"github.com/0shima_x/go-cart/data_loader"
	"github.com/0shima_x/go-cart/splitter"
	"github.com/0shima_x/go-cart/tree"
)

func main() {
	path := "./data.csv"
	datas := data_loader.Load(path)
	fmt.Printf("Data: %v\n", datas)

	X, y := data_loader.SplitIntoFeaturesLabes(datas, 2)
	fmt.Printf("X: %v\n", X)
	fmt.Printf("y: %v\n", y)

	setX1, setX2, setY1, setY2 := splitter.Split(X, y, 0, "google")
	fmt.Printf("setX1: %v\n", setX1)
	fmt.Printf("setX2: %v\n", setX2)
	fmt.Printf("setY1: %v\n", setY1)
	fmt.Printf("setY2: %v\n", setY2)

	gini := criterion.GiniImpurity(y)
	ent := criterion.Entropy(y)
	fmt.Printf("gini: %f\n", gini)
	fmt.Printf("ent: %f\n", ent)

	model := tree.DecisionTreeClassifier{Criterion: "gini"}
	model.Build(X, y)
}
