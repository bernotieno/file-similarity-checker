package main

import (
	"fmt"
	"log"
	"os"
	"similaritychecker/checker"
)

func main() {
	// Default values
	format := "text"               // default format
	directory := ""                // default: current working directory
	outputFile := "report_output" // base file name

	// Read command-line arguments
	if len(os.Args) > 1 {
		format = os.Args[1]
	}
	if len(os.Args) > 2 {
		directory = os.Args[2]
	}

	// Decide file extension based on format
	var fileExt string
	switch format {
	case "html":
		fileExt = ".html"
	case "pdf":
		fileExt = ".pdf"
	default:
		format = "text"
		fileExt = ".txt"
	}

	outputFile += fileExt

	// Initialize checker
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

	err = checker.WriteResultsToFile(results, outputFile, format)
	if err != nil {
		log.Fatalf("Error writing results: %v", err)
	}

	fmt.Printf("File similarity report generated at %s\n", outputFile)
	fmt.Printf("Total comparisons: %d\n", len(results))
}
