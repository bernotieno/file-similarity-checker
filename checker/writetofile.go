package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// WriteResultsToFile saves comparison results
func WriteResultsToFile(results []SimilarityResult, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	header := fmt.Sprintf("%-30s %-30s %-15s %-15s\n", "File 1", "File 2", "Similarity %", "Category")
	writer.WriteString(header)
	writer.WriteString(strings.Repeat("-", 90) + "\n")

	for _, result := range results {
		line := fmt.Sprintf("%-30s %-30s %-15.2f %-15s\n",
			result.File1, result.File2, result.Similarity, result.Category)
		writer.WriteString(line)
	}

	return nil
}
