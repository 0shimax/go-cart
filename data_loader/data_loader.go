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
