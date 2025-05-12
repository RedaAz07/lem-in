package main

import (
	"fmt"
	"os"

	graph "lem-in/bfs"
	"lem-in/parsing"
	"lem-in/printage"
	"lem-in/utils"
)

func main() {
	colony := parsing.Parsing()
	if colony == nil {
		os.Exit(1)
	}
	// fmt.Println(colony)

	paths := graph.FindPaths(colony)
	if len(paths) == 0 {
		fmt.Println("No valid paths found.")
		os.Exit(1)
	}

	fmt.Println(string(parsing.File) + "\n")

	if (len(paths))/(utils.Ants) <= (utils.Ants) {
		utils.Filter = graph.FindDisjointPaths(paths, colony)
	} else {
		utils.Filter = graph.FindDisjointPaths2(paths, colony)
	}

	printage.Sendants(colony)
}
