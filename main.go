package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SimilarityResult stores comparison results between two files
type SimilarityResult struct {
	File1       string
	File2       string
	Similarity  float64
	Category    string
}

// CodeSimilarityChecker manages file comparisons
type CodeSimilarityChecker struct {
	directory string
	files     []string
}

// New creates a new CodeSimilarityChecker
func New(directory string) (*CodeSimilarityChecker, error) {
	var err error
	// If no directory is specified, use current directory
	if directory == "" {
		directory, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	} else {
		// Expand to full path if a relative path is given
		directory, err = filepath.Abs(directory)
		if err != nil {
			return nil, err
		}
	}

	checker := &CodeSimilarityChecker{directory: directory}
	err = checker.findFiles()
	if err != nil {
		return nil, err
	}
	return checker, nil
}

// findFiles identifies files in the directory
func (cc *CodeSimilarityChecker) findFiles() error {
	extensions := []string{
		// Code files
		".go", ".py", ".js", ".cpp", ".java", ".rs", ".c", ".rb", ".html", ".css", ".php", 
		// Text and data files
		".txt", ".csv", ".json", ".xml", ".md", 
		// Image files
		".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp", 
		// Document files
		".pdf", ".doc", ".docx", ".xls", ".xlsx",
	}
	
	err := filepath.Walk(cc.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if info.IsDir() {
			return nil
		}
		
		// Check file extensions
		for _, ext := range extensions {
			if filepath.Ext(path) == ext {
				cc.files = append(cc.files, path)
				break
			}
		}
		
		return nil
	})
	
	if len(cc.files) < 2 {
		return fmt.Errorf("need at least two files to compare")
	}
	
	// Limit to 50 files to prevent excessive comparisons
	if len(cc.files) > 50 {
		cc.files = cc.files[:50]
	}
	
	return err
}

// readFileContent reads entire file content
func readFileContent(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// calculateSimilarity computes similarity between two files
func calculateSimilarity(content1, content2 string) float64 {
	// Reduce content for binary files like images
	if len(content1) > 10000 {
		content1 = content1[:10000]
	}
	if len(content2) > 10000 {
		content2 = content2[:10000]
	}
	
	lines1 := strings.Split(content1, "\n")
	lines2 := strings.Split(content2, "\n")
	
	matchingLines := 0
	for _, line1 := range lines1 {
		for _, line2 := range lines2 {
			if strings.TrimSpace(line1) == strings.TrimSpace(line2) {
				matchingLines++
				break
			}
		}
	}
	
	totalLines := len(lines1) + len(lines2)
	if totalLines == 0 {
		return 0
	}
	
	return (2.0 * float64(matchingLines)) / float64(totalLines) * 100
}

// categorizeSimilarity determines similarity category
func categorizeSimilarity(similarityScore float64) string {
	switch {
	case similarityScore > 70:
		return "Very Similar"
	case similarityScore > 30:
		return "Similar"
	default:
		return "Not Similar"
	}
}

// compareFiles performs comparisons and generates results
func (cc *CodeSimilarityChecker) compareFiles() ([]SimilarityResult, error) {
	var results []SimilarityResult
	
	for i := 0; i < len(cc.files); i++ {
		for j := i + 1; j < len(cc.files); j++ {
			content1, err1 := readFileContent(cc.files[i])
			content2, err2 := readFileContent(cc.files[j])
			
			if err1 != nil || err2 != nil {
				continue // Skip files that can't be read
			}
			
			similarityScore := calculateSimilarity(content1, content2)
			category := categorizeSimilarity(similarityScore)
			
			result := SimilarityResult{
				File1:       filepath.Base(cc.files[i]),
				File2:       filepath.Base(cc.files[j]),
				Similarity:  similarityScore,
				Category:    category,
			}
			results = append(results, result)
		}
	}
	
	return results, nil
}

// writeResultsToFile saves comparison results
func writeResultsToFile(results []SimilarityResult, outputFile string) error {
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

func main() {
	// Check if a custom directory is specified
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}
	
	outputFile := "file_similarity_report.txt"
	
	checker, err := New(directory)
	if err != nil {
		log.Fatalf("Error initializing: %v", err)
	}
	
	fmt.Printf("Analyzing files in directory: %s\n", checker.directory)
	fmt.Printf("Total files found: %d\n", len(checker.files))
	
	results, err := checker.compareFiles()
	if err != nil {
		log.Fatalf("Error comparing files: %v", err)
	}
	
	err = writeResultsToFile(results, outputFile)
	if err != nil {
		log.Fatalf("Error writing results: %v", err)
	}
	
	fmt.Printf("File similarity report generated at %s\n", outputFile)
	fmt.Printf("Total comparisons: %d\n", len(results))
}