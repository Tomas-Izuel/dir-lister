package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//Flag declaration
	flagPattern := flag.String("p", "", "Filter by file path")
	flagAll := flag.Bool("a", false, "List all files")
	flagNumberRecords := flag.Int("n", 10, "Number of records to display")
	flagOrderByTime := flag.Bool("t", false, "Order by time")
	flagOrderBySize := flag.Bool("s", false, "Order by size")
	flagOrderByName := flag.Bool("N", false, "Order by name")

	//Parse the flags
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	memDirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, dir := range memDirs {
		fmt.Println(dir.Name())
	}

	fmt.Println("Pattern:", *flagPattern)
	fmt.Println("All:", *flagAll)
	fmt.Println("Number of records:", *flagNumberRecords)
	fmt.Println("Order by time:", *flagOrderByTime)
	fmt.Println("Order by size:", *flagOrderBySize)
	fmt.Println("Order by name:", *flagOrderByName)
}
