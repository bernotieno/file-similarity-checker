package main

import (
	"fmt"
	"log"
	"os"
	"similaritychecker/checker"
)

func main() {
	// Check if a custom directory is specified
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}

	outputFile := "file_similarity_report.txt"

	checkers, err := checker.New(directory)
	if err != nil {
		log.Fatalf("Error initializing: %v", err)
	}

	fmt.Printf("Analyzing files in directory: %s\n", checkers.Directory())
	fmt.Printf("Total files found: %d\n", len(checkers.Files()))

	results, err := checkers.CompareFiles()
	if err != nil {
		log.Fatalf("Error comparing files: %v", err)
	}

	err = checker.WriteResultsToFile(results, outputFile)
	if err != nil {
		log.Fatalf("Error writing results: %v", err)
	}

	fmt.Printf("File similarity report generated at %s\n", outputFile)
	fmt.Printf("Total comparisons: %d\n", len(results))
}
