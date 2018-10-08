package splitter

import "reflect"

type Predicate func(interface{}, interface{}) bool

func predicateEq(a, b interface{}) bool {
	return a == b
}

func predicateGte(a, b interface{}) bool {
	switch a.(type) {
	case float64:
		a_ := a.(float64)
		b_ := b.(float64)
		return a_ >= b_
	case int:
		a_ := a.(int)
		b_ := b.(int)
		return a_ >= b_
	}
	return false
}

func Split(features [][]interface{}, labels []interface{}, col int, value interface{}) ([][]interface{}, [][]interface{}, []interface{}, []interface{}) {
	// var res Node
	predicate := getPredicte(value)
	var setMatchFeatures [][]interface{}
	var setUnMatchFeatures [][]interface{}

	var setMatchLbels []interface{}
	var setUnMatchLbels []interface{}

	for _, item := range features {
		if predicate(item[col], value) {
			setMatchFeatures = append(setMatchFeatures, item)
			setMatchLbels = append(setMatchLbels, item)
		} else {
			setUnMatchFeatures = append(setUnMatchFeatures, item)
			setUnMatchLbels = append(setUnMatchLbels, item)
		}
	}
	return setMatchFeatures, setUnMatchFeatures, setMatchLbels, setUnMatchLbels
}

func getPredicte(value interface{}) Predicate {
	var predicate Predicate
	if reflect.TypeOf(value).String() == "int" || reflect.TypeOf(value).String() == "float64" {
		predicate = predicateGte
	} else {
		predicate = predicateEq
	}
	return predicate
}
