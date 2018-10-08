package data_loader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Load(path string) [][]interface{} {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var datas [][]interface{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		intrfcLine := make([]interface{}, len(line))
		for i, v := range line {
			intrfcLine[i] = v
		}
		datas = append(datas, intrfcLine)
	}
	return datas
}

func SplitIntoFeaturesLabes(rows [][]interface{}, targetCol int) ([][]interface{}, []interface{}) {
	var X [][]interface{}
	var y []interface{}

	lenRow := len(rows[0])
	for _, row := range rows {
		y = append(y, row[targetCol])

		if lenRow-1 == targetCol {
			feature := row[:targetCol]
			X = append(X, feature)
		} else if targetCol == 0 {
			feature := row[1:]
			X = append(X, feature)
		} else {
			slice1 := row[:targetCol]
			slice2 := row[targetCol+1:]
			slice1 = append(slice1, slice2...)
			X = append(X, slice1)
		}
	}
	return X, y
}
