package main

import (
	"github.com/nounoursheureux/dotrender"
	"log"
)

func main() {
	var output, err = dotrender.RenderFileToString("graph.dot")
	if err != nil {
		log.Fatal(err)
	} else {
		println(output)
	}
}
