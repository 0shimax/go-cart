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

func Split(rows [][]interface{}, col int, value interface{}) ([][]interface{}, [][]interface{}) {
	// var res Node
	predicate := getPredicte(value)
	var setMatch [][]interface{}
	var setUnMatch [][]interface{}

	for _, item := range rows {
		if predicate(item[col], value) {
			setMatch = append(setMatch, item)
		} else {
			setUnMatch = append(setUnMatch, item)
		}
	}
	return setMatch, setUnMatch
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

func uniqueCounts(targets []string) map[string]int {
	var results map[string]int
	for _, target := range targets {
		if _, ok := results[target]; ok {
			results[target] += 1
		} else {
			results[target] = 1
		}
	}
	return results
}
