package tree

import (
	"fmt"
	"log"
	"reflect"

	"github.com/0shima_x/go-cart/criterion"
	"github.com/0shima_x/go-cart/node"
	"github.com/0shima_x/go-cart/splitter"
)

type DecisionTreeClassifier struct {
	Criterion       string
	MaxDepth        int
	MinSamplesSplit int
	NClasses        int
	NFeatures       int
	NFeature        int
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
	var LeftbestFeatureSets [][]interface{}
	var LeftbestLabelSets []interface{}
	var RightbestFeatureSets [][]interface{}
	var RightbestLabelSets []interface{}

	for split_col := 0; split_col < tree.NFeature; split_col++ {
		columnValues := make(map[interface{}]int)
		for _, feature := range features {
			columnValues[feature[split_col]] = 1
		}
		for split_val := range columnValues {
			setX1, setX2, setY1, setY2 := splitter.Split(features, labels, split_col, split_val)
			p := float64(len(setX1)) / float64(tree.NFeatures)
			gain := currentScore - p*metric(setY1) - (1-p)*metric(setY2)
			if gain > bestGain && len(setY1) > 0 && len(setY2) > 0 {
				bestGain = gain
				bestCriteria = map[int]interface{}{split_col: split_val}

				LeftbestFeatureSets = setX1
				RightbestFeatureSets = setX2

				LeftbestLabelSets = setY1
				RightbestLabelSets = setY2
			}
		}
	}

	if bestGain > 0 {
		leftBranch := tree.Build(LeftbestFeatureSets, LeftbestLabelSets)
		rightBranch := tree.Build(RightbestFeatureSets, RightbestLabelSets)
		splitCol := reflect.ValueOf(bestCriteria).MapKeys()[0].Interface().(int)
		return node.Node{SplitCol: splitCol, SplitValue: bestCriteria[splitCol], Left: leftBranch, Right: rightBranch, NodeType: node.BRANCH}
	} else {
		return node.Node{Results: criterion.UniqueCounts(labels), NodeType: node.LEAF}
	}
}

func (tree *DecisionTreeClassifier) Print(trainedModel node.Node, indent string) {
	if trainedModel.NodeType == node.LEAF {
		fmt.Println(trainedModel.Results)
	} else {
		fmt.Println(fmt.Sprint(trainedModel.SplitCol) + ":" + fmt.Sprint(trainedModel.SplitValue) + "?")

		fmt.Print(indent + "T-> ")
		tree.Print(trainedModel.Left.(node.Node), indent+"  ")
		fmt.Print(indent + "F-> ")
		tree.Print(trainedModel.Right.(node.Node), indent+"  ")
	}
}
