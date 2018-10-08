package main

import (
	"fmt"

	"github.com/0shima_x/go-cart/data_loader"
	"github.com/0shima_x/go-cart/splitter"
)

func main() {
	path := "./data.csv"
	datas := data_loader.Load(path)
	set1, set2 := splitter.Split(datas, 2, "yes")

	fmt.Printf("%v", datas)
	// fmt.Printf("%v", set1)
	// fmt.Printf("%v", set2)
}
