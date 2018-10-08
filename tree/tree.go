package tree

import (
	"fmt"
	"log"
	"reflect"

	"github.com/0shima_x/go-cart/criterion"
	"github.com/0shima_x/go-cart/node"
	"github.com/0shima_x/go-cart/splitter"
)

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

type DecisionTreeClassifier struct {
	Criterion       string
	MaxDepth        int
	MinSamplesSplit int
	NClasses        int
	NFeatures       int
	NFeature        int
	Root            Link
}

func (tree *DecisionTreeClassifier) Build(features [][]interface{}, labels []interface{}) node.Node {
	if len(features) == 0 {
		err := fmt.Errorf("error occured: %s", "features sise must be more than 1.")
		log.Fatal(err)
	}
	tree.NFeature = len(features[0])
	tree.NFeatures = len(features)

	var metric criterion.Metric
	if tree.Criterion == "Gini" || tree.Criterion == "gini" || tree.Criterion == "GINI" {
		metric = criterion.GiniImpurity
	} else if tree.Criterion == "Entropy" || tree.Criterion == "entropy" || tree.Criterion == "ENTROPY" {
		metric = criterion.Entropy
	}
	currentScore := metric(labels)

	var bestGain float64
	bestCriteria := make(map[int]interface{})
	var bestFeatureSets [][][]interface{}
	var bestLabelSets [][]interface{}

	for col := 0; col < tree.NFeature; col++ {
		columnValues := make(map[interface{}]int)
		for _, feature := range features {
			columnValues[feature[col]] = 1
		}
		for val := range columnValues {
			setX1, setX2, setY1, setY2 := splitter.Split(features, labels, col, val)
			p := float64(len(setX1)) / float64(tree.NFeatures)
			gain := currentScore - p*metric(setY1) - (1-p)*metric(setY2)
			if gain > bestGain && len(setY1) > 0 && len(setY2) > 0 {
				bestGain = gain

				tmpCriteria := map[int]interface{}{col: val}
				bestCriteria = tmpCriteria

				var tmpSets [][][]interface{}
				tmpSets = append(tmpSets, setX1)
				tmpSets = append(tmpSets, setX2)
				bestFeatureSets = tmpSets

				var tmpLSets [][]interface{}
				tmpLSets = append(tmpLSets, setY1)
				tmpLSets = append(tmpLSets, setY2)
				bestLabelSets = tmpLSets
			}
		}

		if bestGain > 0 {
			leftBranch := tree.Build(bestFeatureSets[0], bestLabelSets[0])
			rightBranch := tree.Build(bestFeatureSets[1], bestLabelSets[1])
			criteriaCol := reflect.ValueOf(bestCriteria).MapKeys()[0].Interface().(int)
			return node.Node{Col: criteriaCol, Value: bestCriteria[criteriaCol], Left: leftBranch, Right: rightBranch}
		} else {
			return node.Node{Results: criterion.UniqueCounts(labels)}
		}
	}
	return node.Node{}
}
