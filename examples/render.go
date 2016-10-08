package main

import (
	"github.com/nounoursheureux/dotrender"
	"log"
)

func main() {
	var err = dotrender.RenderFile("graph.dot", "graph.png")
	if err != nil {
		log.Fatal(err)
	}
}
